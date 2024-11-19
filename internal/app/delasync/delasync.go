// Модуль delasync отвечает за асинхронное удаление данных из базы.
package delasync

import (
	"context"
	"log/slog"
	"time"
)

// Содержит информацию о том, что нужно удалить: сама сущность и кому пренадлежит.
type DelPair struct {
	UserID string
	Item   string
}

// Функция запускает горутины, которые будут принимать информацию об удаляемых объектах.
func InitWorker(ch <-chan DelPair, delFunc func(context.Context, []DelPair) error) {
	go func(in <-chan DelPair) {
		var buf []DelPair // сюда будем складывать накопленные

		rf := func(in <-chan DelPair) (DelPair, bool, bool) {
			select {
			case val, opened := <-in:
				return val, true, opened
			case <-time.After(5 * time.Second):
				return DelPair{}, false, true
			}
		}

		for {
			val, found, opened := rf(in)
			if found {
				buf = append(buf, val)
				if len(buf) > 2 {
					err := delFunc(context.Background(), buf)
					if err != nil {
						slog.Error("Error when execute remove function", "err", err)
					}

					buf = nil
				}
			} else {
				if len(buf) > 0 {
					err := delFunc(context.Background(), buf)
					if err != nil {
						slog.Error("Error when execute remove function", "err", err)
					}

					buf = nil
				}
			}
			if !opened {
				break
			}
		}
	}(ch)
}
