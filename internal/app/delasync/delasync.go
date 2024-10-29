// Модуль delasync отвечает за асинхронное удаление данных из базы.
package delasync

import (
	"context"
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

		rf := func(in <-chan DelPair) (DelPair, bool) {
			select {
			case val := <-in:
				return val, true
			case <-time.After(5 * time.Second):
				return DelPair{}, false
			}
		}

		for {
			val, found := rf(in)
			if found {
				buf = append(buf, val)
				if len(buf) > 2 {
					delFunc(context.Background(), buf)
					buf = nil
				}
			} else {
				if len(buf) > 0 {
					delFunc(context.Background(), buf)
					buf = nil
				}
			}
		}
	}(ch)
}
