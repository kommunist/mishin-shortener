package middleware

import (
	"compress/gzip"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

// Структура компрессора ответа.
type compressWriter struct {
	w  http.ResponseWriter
	zw *gzip.Writer
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

// Получение хедеров.
func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

// Запись результата.
func (c *compressWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

// Запись хедеров.
func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.w.Header().Set("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

// Закрытие записи.
func (c *compressWriter) Close() error {
	return c.zw.Close()
}

// Структура компрессора запроса.
type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

// Чтение запроса.
func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

// Закрытие чтения.
func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}

// Мидлварь компрессии/декомпрессии запроса и ответа
func Gzip(h http.Handler) http.Handler {
	compressFn := func(w http.ResponseWriter, r *http.Request) {
		ow := w // сохранил оригинальный writer

		// Если клиент поддерживает шифрование gzip, то подменяем на свой writer
		acceptEncoding := r.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			cw := newCompressWriter(w)
			ow = cw
			defer cw.Close()
		}

		// Проверяем, а не зашифрован ли контент
		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")

		if sendsGzip { // А если зашфирован, то вставляем gzip reader "между" хендлером и body
			cr, err := newCompressReader(r.Body)
			if err != nil {
				slog.Error("Error when compress read", "err", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = cr
			defer cr.Close()
		}

		h.ServeHTTP(ow, r)
	}
	return http.HandlerFunc(compressFn)

}
