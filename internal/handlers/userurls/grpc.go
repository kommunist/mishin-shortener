package userurls

import (
	"log/slog"
	pb "mishin-shortener/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultUserID = "ququ" // пока не придумал, что делать с userID

func (h *Handler) CallGRPC(ctx context.Context, in *pb.UserUrlsRequest) (*pb.UserUrlsResponse, error) {
	data, err := h.Perform(ctx, defaultUserID)
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
