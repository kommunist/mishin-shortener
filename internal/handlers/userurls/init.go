package userurls

import (
	"context"
	"mishin-shortener/internal/config"
	pb "mishin-shortener/proto"
)

// Интерфейс доступа к базе
type ByUserIDGetter interface {
	GetByUserID(context.Context, string) (map[string]string, error)
}

// Структура хендлера
type Handler struct {
	storage ByUserIDGetter
	setting config.MainConfig

	pb.UnimplementedUserUrlsServer
}

type responseItem struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage ByUserIDGetter) Handler {
	return Handler{storage: storage, setting: setting}
}
