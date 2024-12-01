package ping

import (
	"context"
	"log/slog"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CallGRPC(ctx context.Context) (*pb.PingResponse, error) {
	err := h.Perform(ctx)
	if err != nil {
		slog.Error("Error when get stats from db", "err", err)
		return nil, status.Error(codes.Unknown, "Error when call service")
	}
	return &pb.PingResponse{}, nil
}
