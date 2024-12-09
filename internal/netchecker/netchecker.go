package netchecker

import (
	"log/slog"
	"net"
)

// Основная структура чекера
type Handler struct {
	subNet *net.IPNet
}

// кКонструктор чекера
func Make(inp string) (Handler, error) {
	_, subNet, err := net.ParseCIDR(inp)

	if err != nil {
		slog.Error("Error when parse subNet")
		return Handler{}, err
	}

	return Handler{subNet: subNet}, nil
}

// Метод проверки
func (h *Handler) Contains(ipAddr string) bool {
	parsed := net.ParseIP(ipAddr)
	return h.subNet.Contains(parsed)
}
