package stats

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/netchecker"
	pb "mishin-shortener/proto"
)

// Интерфейс доступа к базе
type StatsGetter interface {
	GetStats(context.Context) (int, int, error)
}

// Структура хендлера
type Handler struct {
	storage    StatsGetter
	netChecker netchecker.Handler

	pb.UnimplementedStatsServer
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage StatsGetter) (Handler, error) {
	netchecker, err := netchecker.Make(setting.TrustedSubnet)

	if err != nil {
		slog.Error("Error when set netchecker")
		return Handler{}, err
	}

	return Handler{
		storage:    storage,
		netChecker: netchecker,
	}, nil
}
