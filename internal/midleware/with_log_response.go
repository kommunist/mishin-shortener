package middleware

import "net/http"

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter

	responseData *responseData
}

// Запись тела ответа. Сохраняет размер ответа в структуру лога ответа
func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

// Запись хедеров ответа. Сохраняет статус ответа в структуру лога ответа
func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}
