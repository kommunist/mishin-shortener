package app

import (
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	// данный тест проверяет непосредственно запуск сервера и
	// что вся предварительная работа с роутерами, базой и конфигом проходит без проблем
	// и сервер запускается
	t.Run("happy_path_on_start_full", func(t *testing.T) {
		h := Make()
		go func() {
			time.Sleep(2 * time.Second)
			h.API.Stop()
		}()
		h.Call()
	})

}
