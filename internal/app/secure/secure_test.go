package secure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func EncryptDecryptTest(t *testing.T) {
	t.Run("original_string_after_ancrypt_and_decrypt", func(t *testing.T) {
		original := "Lenin is a live!"

		encrypted, _ := Encrypt(original)
		assert.NotEqual(t, original, encrypted)

		decrypted, _ := Decrypt(encrypted)
		assert.Equal(t, original, decrypted)
	})

}
