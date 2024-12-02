package createjsonbatch

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CallGRPC(ctx context.Context, in *pb.CreateBatchRequest) (*pb.CreateBatchResponse, error) {
	var userID string
	if ctx.Value(secure.UserIDKey) == nil {
		return nil, status.Error(codes.Unknown, "Error with auth")
	} else {
		userID = ctx.Value(secure.UserIDKey).(string)
	}

	input := make([]requestBatchItem, 0, len(in.List))
	response := make([]*pb.CreateBatchResponseItem, 0, len(in.List))

	for _, item := range in.List {
		input = append(
			input, requestBatchItem{CorrelationID: item.CorrelationId, OriginalURL: item.OriginalUrl},
		)
	}

	output, err := h.Perform(ctx, input, userID)
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
