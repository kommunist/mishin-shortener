package pgstorage

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"mishin-shortener/internal/app/exsist"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

func (d *Driver) Push(short string, original string) error {
	err := insert(short, original, false, d.driver)

	if err != nil {
		slog.Error("When push to db error", "err", err)

		if errStruct, ok := err.(*pq.Error); ok && errStruct.Code == pgerrcode.UniqueViolation {
			return exsist.NewExistError(err) // если уже существует такая запись, то возвращаем ошибку
		}

		return err
	}
	return nil
}

func (d *Driver) PushBatch(list *map[string]string) error {
	tx, err := d.driver.Begin()
	if err != nil {
		slog.Error("When open transaction error", "err", err)
		return err
	}

	for k, v := range *list {
		err := insert(k, v, true, tx) // в случае инстерта батчами будем с форсом
		if err != nil {
			slog.Error("When batch insert error", "err", err)
			tx.Rollback()
			return err
		}

	}
	tx.Commit()
	return nil
}

func insert(short string, original string, force bool, d interface {
	Exec(string, ...any) (sql.Result, error)
}) error {
	query := "INSERT INTO short_urls (short, original) VALUES ($1, $2)"
	if force {
		query += " ON CONFLICT (short) DO UPDATE SET original = excluded.original"
	}
	_, err := d.Exec(query, short, original)

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
