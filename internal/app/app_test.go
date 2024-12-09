package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	// данный тест проверяет непосредственно запуск сервера и
	// что вся предварительная работа с роутерами, базой и конфигом проходит без проблем
	// и сервер запускается
	t.Run("happy_path_on_start_full", func(t *testing.T) {
		t.Setenv("ENABLE_PROFILE", "true")

		h, err := Make()
		assert.NoError(t, err)
		go func() {
			time.Sleep(5 * time.Second)
			h.stop()
		}()
		err = h.Call()
		assert.NoError(t, err)
	})

	t.Run("happy_path_on_start_with_tls_full", func(t *testing.T) {
		t.Setenv("ENABLE_HTTPS", "true")
		t.Setenv("ENABLE_PROFILE", "true")
		t.Setenv("CERT_PATH", "../../certs/MyCertificate.crt")
		t.Setenv("CERT_KEY_PATH", "../../certs/MyKey.key")

		h, err := Make()
		assert.NoError(t, err)
		go func() {
			time.Sleep(5 * time.Second)
			h.stop()
		}()
		err = h.Call()
		assert.NoError(t, err)
	})

}
