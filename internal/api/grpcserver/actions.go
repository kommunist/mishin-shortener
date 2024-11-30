package grpcserver

import (
	"context"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *GRPCHandler) GetStats(ctx context.Context, in *pb.GetStatsRequest) (*pb.GetStatsResponse, error) {
	resp, err := h.stats.CallGRPC(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Error when call service", err)
	}
	return resp, nil
}
