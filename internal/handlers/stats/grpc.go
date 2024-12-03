package stats

import (
	"context"
	"log/slog"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CallGRPC(ctx context.Context) (*pb.GetStatsResponse, error) {
	users, urls, err := h.Perform(ctx)
	if err != nil {
		slog.Error("Error when get stats from db", "err", err)
		return nil, status.Error(codes.Unknown, "Error when call service")
	}
	return &pb.GetStatsResponse{CountUsers: uint32(users), CountUrls: uint32(urls)}, nil
}
