package pgstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"mishin-shortener/internal/app/deleted"
	"mishin-shortener/internal/app/exsist"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

func (d *Driver) Push(ctx context.Context, short string, original string, userID string) error {
	err := insert(ctx, short, original, userID, false, d.driver)
	if err != nil {
		slog.Error("When push to db error", "err", err)

		if errStruct, ok := err.(*pq.Error); ok && errStruct.Code == pgerrcode.UniqueViolation {
			return exsist.NewExistError(err) // если уже существует такая запись, то возвращаем ошибку
		}

		return err
	}
	return nil
}

func (d *Driver) PushBatch(ctx context.Context, list *map[string]string, userID string) error {
	tx, err := d.driver.Begin()
	if err != nil {
		slog.Error("When open transaction error", "err", err)
		return err
	}

	for k, v := range *list {
		err := insert(ctx, k, v, userID, true, tx) // в случае инстерта батчами будем с форсом
		if err != nil {
			slog.Error("When batch insert error", "err", err)
			tx.Rollback()
			return err
		}

	}
	tx.Commit()
	return nil
}

// вот тут правильнее было бы использовать какой-то библиотечый интерфейс в качестве аргумента,
// но я такой не нашел
func insert(ctx context.Context, short string, original string, userID string, force bool, d interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}) error {
	query := "INSERT INTO short_urls (short, original, user_id) VALUES ($1, $2, $3)"
	if force {
		query += " ON CONFLICT (short) DO UPDATE SET original = excluded.original"
	}
	_, err := d.ExecContext(ctx, query, short, original, userID)

	if err != nil {
		slog.Error("When insert to db error", "err", err)
		return err
	}
	return nil
}

func (d *Driver) Get(ctx context.Context, short string) (string, error) {
	var result string
	var dR bool

	row := d.driver.QueryRowContext(ctx, "SELECT original, deleted FROM short_urls where short = $1 LIMIT 1", short)

	err := row.Scan(&result, &dR)
	if err != nil {
		slog.Error("When scan data from select", "err", err)
		return "", err
	}

	if result == "" {
		return "", errors.New("not found")
	}

	if dR {
		return "", deleted.NewDeletedError(nil)
	}

	return result, nil

}

func (d *Driver) GetByUserID(ctx context.Context, userID string) (map[string]string, error) {
	rows, err := d.driver.QueryContext(ctx, "SELECT short, original FROM short_urls where user_id = $1", userID)
	if err != nil {
		slog.Error("When select data from db", "err", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]string)

	for rows.Next() {
		var short string
		var original string

		err := rows.Scan(&short, &original)
		if err != nil {
			slog.Error("When scan data from select", "err", err)
			return nil, err
		}
		result[short] = original
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Driver) DeleteByUserID(ctx context.Context, list [][2]string) error {
	start := "UPDATE short_urls set deleted = true where "
	var cond []string
	cond = append(cond, start)

	for i, v := range list {
		if i != 0 {
			cond = append(cond, " or ")
		}
		cond = append(cond, fmt.Sprintf("user_id = '%s' and short = '%s'", v[0], v[1]))
	}
	resCond := strings.Join(cond, "")
	slog.Info("Execute cond", "cond", resCond)
	_, err := d.driver.Exec(resCond)
	if err != nil {
		slog.Error("Error from DeleteByUserID", "err", err)
		return err
	}
	return nil
}

// для проверки, что живо соединение
func (d *Driver) Ping(ctx context.Context) error {
	return d.driver.PingContext(ctx)
}

// заглушка для поддержки интерфейса
func (d *Driver) Finish() error {
	return nil
}
