package ping

import (
	"context"
	"log/slog"
	pb "mishin-shortener/proto"
)

func (h *Handler) CallGRPC(ctx context.Context) (*pb.PingResponse, error) {
	err := h.Perform(ctx)
	if err != nil {
		slog.Error("Error when get stats from db")
		return nil, err
	}
	return &pb.PingResponse{}, nil
}
