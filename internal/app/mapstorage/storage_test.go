package mapstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	db := Storage{}

	t.Run("simple_push_data_to_database", func(t *testing.T) {
		db.Push("key", "value")
		value, _ := db.Get("key")
		assert.Equal(t, value, "value")
	})
}

func TestPushBatch(t *testing.T) {
	db := Storage{}

	t.Run("simple_push_batch_data_to_database", func(t *testing.T) {
		data := make(map[string]string)
		data["key"] = "value"
		data["biba"] = "boba"

		db.PushBatch(&data)

		var v string
		v, _ = db.Get("key")
		assert.Equal(t, "value", v)

		v, _ = db.Get("biba")
		assert.Equal(t, "boba", v)
	})
}

func TestGet(t *testing.T) {
	db := Storage{"key": "value"}

	t.Run("simple_get_data_from_database", func(t *testing.T) {
		value, err := db.Get("key")
		assert.Equal(t, value, "value")
		assert.Equal(t, err, nil)
	})

	t.Run("simple_get_data_from_database_when_not_found", func(t *testing.T) {
		value, err := db.Get("another_key")
		assert.Equal(t, value, "")
		assert.EqualError(t, err, "not found")
	})
}
