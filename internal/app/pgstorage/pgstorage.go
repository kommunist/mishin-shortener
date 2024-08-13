package pgstorage

import (
	"context"
	"database/sql"
	"log/slog"
	"mishin-shortener/internal/app/config"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Driver struct {
	driver *sql.DB
}

func Make(c config.MainConfig) *Driver {
	driver, err := sql.Open("postgres", c.DatabaseDSN)

	if err != nil {
		slog.Error("Eror when connect to database", "err", err)
		os.Exit(1)
	}

	slog.Info("Success connect to database")

	result := Driver{driver: driver}
	result.initSchema()

	return &result
}

// тут хорошо было бы, наверное, затащить какой-нибудь мигратор, но пока обойдемся IF NOT EXIST
func (d *Driver) initSchema() {
	// здесь нет внешнего, поэтому нужен отдельный context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := d.driver.ExecContext(
		ctx,
		"CREATE TABLE IF NOT EXISTS short_urls (id SERIAL PRIMARY KEY, short TEXT, original TEXT);")
	if err != nil {
		slog.Error("Eror when create table", "err", err)
		os.Exit(1)
	}

	_, err = d.driver.ExecContext(
		ctx,
		"CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_short_urls_short on short_urls (short);")
	if err != nil {
		slog.Error("Eror when create index", "err", err)
		os.Exit(1)
	}

}
