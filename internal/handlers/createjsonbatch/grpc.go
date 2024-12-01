package createjsonbatch

import (
	"context"
	"log/slog"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultUserID = "default_user_id" // grpc default user id

func (h *Handler) CallGRPC(ctx context.Context, in *pb.CreateBatchRequest) (*pb.CreateBatchResponse, error) {
	input := make([]requestBatchItem, 0, len(in.List))
	response := make([]*pb.CreateBatchResponseItem, 0, len(in.List))

	for _, item := range in.List {
		input = append(
			input, requestBatchItem{CorrelationID: item.CorrelationId, OriginalURL: item.OriginalUrl},
		)
	}

	output, err := h.Perform(ctx, input, defaultUserID)
	if err != nil {
		slog.Error("Error when get stats from db", "err", err)
		return nil, status.Error(codes.Unknown, "Error when call service")
	}

	for _, item := range output {
		response = append(
			response, &pb.CreateBatchResponseItem{CorrelationId: item.CorrelationID, ShortUrl: item.ShortURL},
		)
	}

	return &pb.CreateBatchResponse{List: response}, nil

}
