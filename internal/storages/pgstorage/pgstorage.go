// Модуль pgstorage предоставляет необходимые приложению методы работы с postgres.
package pgstorage

import (
	"context"
	"database/sql"
	"log/slog"
	"mishin-shortener/internal/config"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const initTableQuery = `
	CREATE TABLE IF NOT EXISTS short_urls (
		id SERIAL PRIMARY KEY,
		short     TEXT,
		original  TEXT,
		user_id   TEXT,
		deleted   BOOLEAN DEFAULT false
	);
`

const initIndexQuery = `
  CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_short_urls_short on short_urls (short);
`

// Структура хранилища
type Driver struct {
	driver *sql.DB
}

// Функция создания хранилища
func Make(c config.MainConfig) *Driver {
	driver, err := sql.Open("postgres", c.DatabaseDSN)

	if err != nil {
		slog.Error("Eror when connect to database", "err", err)
		os.Exit(1)
	}

	slog.Info("Success connect to database")

	result := Driver{driver: driver}
	result.initSchema(context.Background())

	return &result
}

// тут хорошо было бы, наверное, затащить какой-нибудь мигратор, но пока обойдемся IF NOT EXIST
func (d *Driver) initSchema(ctx context.Context) {
	// здесь нет внешнего, поэтому нужен отдельный context
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	_, err := d.driver.ExecContext(ctx, initTableQuery)
	if err != nil {
		slog.Error("Eror when create table", "err", err)
		os.Exit(1)
	}
	_, err = d.driver.ExecContext(ctx, initIndexQuery)
	if err != nil {
		slog.Error("Eror when create index", "err", err)
		os.Exit(1)
	}
}
