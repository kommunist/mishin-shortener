package pgstorage

import (
	"context"
	"errors"
	"log/slog"
)

// чтобы поддержать логику "кто последний, тот и прав" как было в mapstorage
// сделаем он conflict update
func (d *Driver) Push(short string, original string) error {
	_, err := d.driver.Exec(
		`
		 INSERT INTO short_urls (short, original) 
		 VALUES ($1, $2) 
		 ON CONFLICT (short) DO UPDATE SET original = excluded.original 
		`, short, original,
	)
	if err != nil {
		slog.Error("When insert to db error", "err", err)
		return err
	}
	return nil
}

func (d *Driver) Get(short string) (string, error) {
	var result string

	row := d.driver.QueryRow("SELECT original FROM short_urls where short = $1 LIMIT 1", short)

	err := row.Scan(&result)
	if err != nil {
		slog.Error("When scan data from select", "err", err)
		return "", err
	}

	if result == "" {
		return "", errors.New("not found")
	}

	return result, nil

}

// для проверки, что живо соединение
func (d *Driver) Ping(ctx context.Context) error {
	return d.driver.PingContext(ctx)
}

// заглушка для поддержки интерфейса
func (d *Driver) Finish() error {
	return nil
}
