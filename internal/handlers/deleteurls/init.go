package deleteurls

import (
	"mishin-shortener/internal/delasync"
	pb "mishin-shortener/proto"
)

// Структура хендлера
type Handler struct {
	DelChan chan delasync.DelPair

	pb.UnimplementedDeleteUrlsServer
}

// Конструктор хендлера
func Make(c chan delasync.DelPair) Handler {
	return Handler{DelChan: c}
}
