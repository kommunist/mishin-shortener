package userurls

import (
	"log/slog"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) UserUrls(ctx context.Context, in *pb.UserUrlsRequest) (*pb.UserUrlsResponse, error) {
	var userID string
	if ctx.Value(secure.UserIDKey) == nil {
		return nil, status.Error(codes.Unknown, "Error with auth")
	} else {
		userID = ctx.Value(secure.UserIDKey).(string)
	}
	data, err := h.Perform(ctx, userID)
	if err != nil {
		slog.Error("Error when perform service")
		return nil, status.Error(codes.Unknown, "Error when call service")
	}

	if len(data) == 0 {
		return nil, status.Error(codes.NotFound, "No data")
	}

	list := make([]*pb.UserUrlsResponseItem, 0, len(data))
	for _, item := range data {
		list = append(list, &pb.UserUrlsResponseItem{Short: item.Short, Original: item.Original})
	}
	return &pb.UserUrlsResponse{List: list}, nil

}
