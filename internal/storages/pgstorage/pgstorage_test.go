package pgstorage

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInitSchema(t *testing.T) {
	exList := []struct {
		name          string
		withErrFirst  bool
		withErrSecond bool
	}{
		{
			name:          "init_schema_happy_path",
			withErrFirst:  false,
			withErrSecond: false,
		},
		{
			name:          "when_error_in_first_request",
			withErrFirst:  true,
			withErrSecond: false,
		},
		{
			name:          "when_error_in_second_request",
			withErrFirst:  false,
			withErrSecond: true,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			stor := Driver{driver: db}

			if ex.withErrFirst {
				mock.ExpectExec(initTableQuery).
					WithoutArgs().
					WillReturnResult(sqlmock.NewResult(0, 0)).
					WillReturnError(errors.New("qq"))
			} else {
				mock.ExpectExec(initTableQuery).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			}
			if !ex.withErrFirst {
				if ex.withErrSecond {
					mock.ExpectExec(initIndexQuery).WithoutArgs().
						WillReturnResult(sqlmock.NewResult(0, 0)).
						WillReturnError(errors.New("qq"))
				} else {
					mock.ExpectExec(initIndexQuery).WithoutArgs().
						WillReturnResult(sqlmock.NewResult(0, 0))
				}
			}

			err = stor.initSchema(context.Background())

			if ex.withErrFirst || ex.withErrSecond {
				assert.EqualError(t, err, errors.New("qq").Error())
			} else {
				assert.NoError(t, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

		})
	}
}
