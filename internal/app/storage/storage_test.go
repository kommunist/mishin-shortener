package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	db := Database{}

	t.Run("simple_push_data_to_database", func(t *testing.T) {
		db.Push("key", "value")
		value, _ := db.Get("key")
		assert.Equal(t, value, "value")
	})
}

func TestGet(t *testing.T) {
	db := Database{"key": "value"}

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
