package deleteurls

import (
	"mishin-shortener/internal/app/delasync"
)

// Структура хендлера
type Handler struct {
	DelChan chan delasync.DelPair
}

// Конструктор хендлера
func Make(c chan delasync.DelPair) Handler {
	return Handler{DelChan: c}
}
