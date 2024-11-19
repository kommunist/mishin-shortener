package pgstorage

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInitSchema(t *testing.T) {
	t.Run("init_schema_happy_path", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		stor := Driver{driver: db}

		mock.ExpectExec(initSchemaQuery).WillReturnResult(sqlmock.NewResult(0, 0))

		stor.initSchema(context.Background())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}
