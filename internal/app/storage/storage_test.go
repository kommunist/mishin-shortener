package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	db := Database{}

	t.Run("simple push test", func(t *testing.T) {
		db.Push("key", "value")
		assert.Equal(t, db["key"], "value")
	})
}

func TestGet(t *testing.T) {
	db := Database{"key": "value"}

	t.Run("simple get test", func(t *testing.T) {
		value := db.Get("key")
		assert.Equal(t, value, "value")
	})
}
