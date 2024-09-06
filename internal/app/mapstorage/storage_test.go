package mapstorage

import (
	"context"
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
