// Модуль delasync отвечает за асинхронное удаление данных из базы.
package delasync

import (
	"context"
	"log/slog"
	"time"
)

// Интерфейс метода для удаления объектов
type Remover interface {
	DeleteByUserID(context.Context, []DelPair) error // слайс пар userID, list
}

// Основная структура
type Handler struct {
	DelChan chan DelPair
	storage Remover
}

// Конструктор основной структуры
func Make(storage Remover) Handler {
	return Handler{
		storage: storage,
		DelChan: make(chan DelPair, 5),
	}
}

// Содержит информацию о том, что нужно удалить: сама сущность и кому пренадлежит.
type DelPair struct {
	UserID string
	Item   string
}

// Функция запускает горутины, которые будут принимать информацию об удаляемых объектах.
func (h *Handler) InitWorker() {
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
					err := h.storage.DeleteByUserID(context.Background(), buf)
					if err != nil {
						slog.Error("Error when execute remove function", "err", err)
					}

					buf = nil
				}
			} else {
				if len(buf) > 0 {
					err := h.storage.DeleteByUserID(context.Background(), buf)
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
	}(h.DelChan)
}

// Метод для остановки
func (h *Handler) Stop() {
	close(h.DelChan)
}
