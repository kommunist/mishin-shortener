package grpcserver

import (
	"context"
	pb "mishin-shortener/proto"
)

// здесь не представлен handler createjson, но он по поведению полностью эквивалентен простому Create

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

func (h *GRPCHandler) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	resp, err := h.redirect.CallGRPC(ctx, in)
	return resp, err
}

func (h *GRPCHandler) UserUrls(ctx context.Context, in *pb.UserUrlsRequest) (*pb.UserUrlsResponse, error) {
	resp, err := h.userUrls.CallGRPC(ctx, in)
	return resp, err
}

func (h *GRPCHandler) DeleteUrls(ctx context.Context, in *pb.DeleteUrlsRequest) (*pb.DeleteUrlsResponse, error) {
	resp, err := h.deleteURLs.CallGRPC(ctx, in)
	return resp, err
}
