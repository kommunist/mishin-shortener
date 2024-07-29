package filestorage

import (
	"mishin-shortener/internal/app/mapstorage"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("get_data_from_cache", func(t *testing.T) {
		db := mapstorage.Make()
		db.Push("short", "original")

		fs := Storage{cache: *db, file: nil}

		value, err := fs.Get("short")
		assert.Equal(t, value, "original")

		require.NoError(t, err)
	})
}
