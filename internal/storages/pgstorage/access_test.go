package pgstorage

import (
	"context"
	"mishin-shortener/internal/delasync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPush(t *testing.T) {
	t.Run("it_correct_push_data", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectExec(
			"INSERT INTO short_urls (short, original, user_id) VALUES ($1, $2, $3)",
		).WithArgs("short", "original", "userID").WillReturnResult(sqlmock.NewResult(1, 1))

		stor := Driver{driver: db}

		stor.Push(context.Background(), "short", "original", "userID")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestPushBatch(t *testing.T) {
	t.Run("it_correct_push_data", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		stor := Driver{driver: db}

		mock.ExpectBegin()
		mock.ExpectExec(
			"INSERT INTO short_urls (short, original, user_id) VALUES ($1, $2, $3) ON CONFLICT (short) DO UPDATE SET original = excluded.original",
		).WithArgs("short", "original", "userID").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		stor.PushBatch(context.Background(), &map[string]string{"short": "original"}, "userID")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("it_correct_select_data", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		stor := Driver{driver: db}

		result := sqlmock.NewRows([]string{"original", "deleted"}).AddRow("original", false)
		mock.ExpectQuery(
			"SELECT original, deleted FROM short_urls where short = $1 LIMIT 1",
		).WithArgs("short").WillReturnRows(result)

		stor.Get(context.Background(), "short")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
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
	t.Run("it_correct_push_data", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		stor := Driver{driver: db}

		mock.ExpectBegin()

		mock.ExpectExec(
			"UPDATE short_urls set deleted = true where user_id = 'userID' and short = 'short'",
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		stor.DeleteByUserID(context.Background(), []delasync.DelPair{{UserID: "userID", Item: "short"}})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
