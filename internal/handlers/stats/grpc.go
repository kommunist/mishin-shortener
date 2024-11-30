package stats

import (
	"context"
	"log/slog"
	pb "mishin-shortener/proto"
)

func (h *Handler) CallGRPC(ctx context.Context) (*pb.GetStatsResponse, error) {
	users, urls, err := h.Perform(ctx)
	if err != nil {
		slog.Error("Error when get stats from db")
		return nil, err
	}
	return &pb.GetStatsResponse{CountUsers: uint32(users), CountUrls: uint32(urls)}, nil
}
