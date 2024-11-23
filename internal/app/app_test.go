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
		h, err := Make()
		assert.NoError(t, err)
		go func() {
			time.Sleep(2 * time.Second)
			h.API.Stop()
		}()
		err = h.Call()
		assert.NoError(t, err)
	})

}
