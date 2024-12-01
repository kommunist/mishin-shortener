package grpcserver

import (
	"context"
	pb "mishin-shortener/proto"
)

func (h *GRPCHandler) GetStats(ctx context.Context, in *pb.GetStatsRequest) (*pb.GetStatsResponse, error) {
	resp, err := h.stats.CallGRPC(ctx)
	return resp, err
}

func (h *GRPCHandler) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	resp, err := h.ping.CallGRPC(ctx)
	return resp, err
}

func (h *GRPCHandler) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	resp, err := h.simpleCreate.CallGRPC(ctx, in)
	return resp, err
}

func (h *GRPCHandler) CreateBatch(ctx context.Context, in *pb.CreateBatchRequest) (*pb.CreateBatchResponse, error) {
	resp, err := h.createJSONBatch.CallGRPC(ctx, in)
	return resp, err
}
