package userurls

import (
	"context"
	"mishin-shortener/internal/app/config"
)

type ByUserIDGetter interface {
	GetByUserID(context.Context, string) (map[string]string, error)
}

type Handler struct {
	storage ByUserIDGetter
	setting config.MainConfig
}

func Make(setting config.MainConfig, storage ByUserIDGetter) Handler {
	return Handler{storage: storage, setting: setting}
}
