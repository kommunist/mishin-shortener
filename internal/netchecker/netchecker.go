package netchecker

import (
	"log/slog"
	"net"
)

type Handler struct {
	subNet *net.IPNet
}

func Make(inp string) (Handler, error) {
	_, subNet, err := net.ParseCIDR(inp)

	if err != nil {
		slog.Error("Error when parse subNet")
		return Handler{}, err
	}

	return Handler{subNet: subNet}, nil
}

func (h *Handler) Contains(ipAddr string) bool {
	parsed := net.ParseIP(ipAddr)
	return h.subNet.Contains(parsed)
}
