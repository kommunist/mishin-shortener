package mapstorage

import (
	"context"
	"mishin-shortener/internal/delasync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	db := Storage{}

	t.Run("simple_push_data_to_database", func(t *testing.T) {
		db.Push(context.Background(), "key", "value", "userID")
		value, _ := db.Get(context.Background(), "key")
		assert.Equal(t, value, "value")
	})
}

func TestPushBatch(t *testing.T) {
	db := Storage{}

	t.Run("simple_push_batch_data_to_database", func(t *testing.T) {
		data := make(map[string]string)
		data["key"] = "value"
		data["biba"] = "boba"

		db.PushBatch(context.Background(), &data, "userID")

		var v string
		v, _ = db.Get(context.Background(), "key")
		assert.Equal(t, "value", v)

		v, _ = db.Get(context.Background(), "biba")
		assert.Equal(t, "boba", v)
	})
}

func TestGet(t *testing.T) {
	db := Storage{"key": "value"}

	t.Run("simple_get_data_from_database", func(t *testing.T) {
		value, err := db.Get(context.Background(), "key")
		assert.Equal(t, value, "value")
		assert.Equal(t, err, nil)
	})

	t.Run("simple_get_data_from_database_when_not_found", func(t *testing.T) {
		value, err := db.Get(context.Background(), "another_key")
		assert.Equal(t, value, "")
		assert.EqualError(t, err, "not found")
	})
}

func TestGetByUserID(t *testing.T) {
	db := Storage{}

	t.Run("empty_test", func(t *testing.T) {
		res, err := db.GetByUserID(context.Background(), "qq")
		assert.Equal(t, map[string]string{}, res)
		assert.NoError(t, err)

	})
}

func TestDeleteByUserID(t *testing.T) {
	db := Storage{}

	t.Run("empty_test", func(t *testing.T) {
		err := db.DeleteByUserID(context.Background(), []delasync.DelPair{})
		assert.NoError(t, err)

	})
}

func TestGetStats(t *testing.T) {
	db := Storage{}

	t.Run("empty_test", func(t *testing.T) {
		users, urls, err := db.GetStats(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 0, users)
		assert.Equal(t, 0, urls)
	})
}
