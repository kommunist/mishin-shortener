package pgstorage

import (
	"context"
	"errors"
	"mishin-shortener/internal/delasync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
	}{
		{
			name:    "it_correct_push_data",
			storErr: nil,
		},
		{
			name:    "receive_another_error_when_push",
			storErr: errors.New("qq"),
		},
		{
			name:    "receive_error_exist",
			storErr: &pq.Error{Code: pgerrcode.UniqueViolation},
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if ex.storErr != nil {
				mock.ExpectExec(
					"INSERT INTO short_urls (short, original, user_id) VALUES ($1, $2, $3)",
				).WithArgs("short", "original", "userID").WillReturnError(ex.storErr)
			} else {
				mock.ExpectExec(
					"INSERT INTO short_urls (short, original, user_id) VALUES ($1, $2, $3)",
				).WithArgs("short", "original", "userID").WillReturnResult(sqlmock.NewResult(1, 1))

			}

			stor := Driver{driver: db}

			err = stor.Push(context.Background(), "short", "original", "userID")
			if ex.storErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, ex.storErr.Error())
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})

	}

}

func TestPushBatch(t *testing.T) {

	exList := []struct {
		name        string
		beginErr    error
		insertErr   error
		commitErr   error
		rollbackErr error
	}{
		{
			name:        "it_correct_push_data",
			beginErr:    nil,
			insertErr:   nil,
			commitErr:   nil,
			rollbackErr: nil,
		},
		{
			name:        "when_begin_error",
			beginErr:    errors.New("qq"),
			insertErr:   nil,
			commitErr:   nil,
			rollbackErr: nil,
		},
		{
			name:        "when_insert_error",
			beginErr:    nil,
			insertErr:   errors.New("qq"),
			commitErr:   nil,
			rollbackErr: nil,
		},
		{
			name:        "when_insert_and_rollback_error",
			beginErr:    nil,
			insertErr:   errors.New("qq"),
			commitErr:   nil,
			rollbackErr: errors.New("qq"),
		},
		{
			name:        "when_commit_error",
			beginErr:    nil,
			insertErr:   nil,
			commitErr:   errors.New("qq"),
			rollbackErr: nil,
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

			if ex.beginErr != nil {
				mock.ExpectBegin().WillReturnError(ex.beginErr)
			} else {
				mock.ExpectBegin()
			}

			if ex.beginErr == nil {
				if ex.insertErr != nil {
					mock.ExpectExec(
						"INSERT INTO short_urls (short, original, user_id) VALUES ($1, $2, $3) ON CONFLICT (short) DO UPDATE SET original = excluded.original",
					).WithArgs("short", "original", "userID").WillReturnError(ex.insertErr)
					if ex.rollbackErr != nil {
						mock.ExpectRollback().WillReturnError(ex.rollbackErr)
					} else {
						mock.ExpectRollback()
					}
				} else {
					mock.ExpectExec(
						"INSERT INTO short_urls (short, original, user_id) VALUES ($1, $2, $3) ON CONFLICT (short) DO UPDATE SET original = excluded.original",
					).WithArgs("short", "original", "userID").WillReturnResult(sqlmock.NewResult(1, 1))

					if ex.commitErr != nil {
						mock.ExpectCommit().WillReturnError(ex.commitErr)
					} else {
						mock.ExpectCommit()
					}

				}

			}

			stor.PushBatch(context.Background(), &map[string]string{"short": "original"}, "userID")

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})

	}

}

func TestGet(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
		rows    *sqlmock.Rows
	}{
		{
			name:    "it_correct_select_data",
			storErr: nil,
			rows:    sqlmock.NewRows([]string{"original", "deleted"}).AddRow("original", false),
		},
		{
			name:    "it_return_any_error",
			storErr: errors.New("ququ"),
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

			if ex.storErr != nil {
				mock.ExpectQuery(
					"SELECT original, deleted FROM short_urls where short = $1 LIMIT 1",
				).WithArgs("short").WillReturnError(ex.storErr)
			} else {
				mock.ExpectQuery(
					"SELECT original, deleted FROM short_urls where short = $1 LIMIT 1",
				).WithArgs("short").WillReturnRows(ex.rows)
			}

			_, err = stor.Get(context.Background(), "short")
			if ex.storErr != nil {
				assert.EqualError(t, err, ex.storErr.Error())
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})

	}

}

func TestGetByUserID(t *testing.T) {
	t.Run("it_correct_select_data", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		stor := Driver{driver: db}

		result := sqlmock.NewRows([]string{"short", "original"}).AddRow("short", "original")
		mock.ExpectQuery(
			"SELECT short, original FROM short_urls where user_id = $1",
		).WithArgs("userID").WillReturnRows(result)

		stor.GetByUserID(context.Background(), "userID")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestDeleteByUserID(t *testing.T) {
	t.Run("it_correct_delete_data", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		stor := Driver{driver: db}

		mock.ExpectBegin()

		mock.ExpectExec(
			"UPDATE short_urls set deleted = true where user_id = 'userID' and short = 'short' or user_id = 'anotherID' and short = 'long'",
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		stor.DeleteByUserID(
			context.Background(), []delasync.DelPair{
				{UserID: "userID", Item: "short"},
				{UserID: "anotherID", Item: "long"},
			},
		)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestGetStats(t *testing.T) {
	exList := []struct {
		name     string
		usersErr error
		urlsErr  error
	}{
		{
			name:     "it_correct_get_stats_from_db",
			usersErr: nil,
			urlsErr:  nil,
		},
		{
			name:     "when_error_in_users_query",
			usersErr: errors.New("ququ"),
			urlsErr:  nil,
		},
		{
			name:     "when_error_in_urls_query",
			usersErr: nil,
			urlsErr:  errors.New("ququ"),
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

			if ex.usersErr != nil {
				mock.ExpectQuery(
					"SELECT count(*) from users;",
				).WithoutArgs().WillReturnError(ex.usersErr)
			} else {
				mock.ExpectQuery(
					"SELECT count(*) from users;",
				).WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

				if ex.urlsErr != nil {
					mock.ExpectQuery(
						"SELECT count(*) from short_urls;",
					).WithoutArgs().WillReturnError(ex.urlsErr)
				} else {
					mock.ExpectQuery(
						"SELECT count(*) from short_urls;",
					).WithoutArgs().WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				}
			}

			stor.GetStats(context.Background())

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})

	}

}
