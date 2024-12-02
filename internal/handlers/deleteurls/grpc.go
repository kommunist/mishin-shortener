package deleteurls

import (
	"context"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CallGRPC(ctx context.Context, in *pb.DeleteUrlsRequest) (*pb.DeleteUrlsResponse, error) {
	var userID string
	if ctx.Value(secure.UserIDKey) == nil {
		return nil, status.Error(codes.Unknown, "Error with auth")
	} else {
		userID = ctx.Value(secure.UserIDKey).(string)
	}

	list := make([]string, 0, len(in.List))
	for _, item := range in.List {
		list = append(list, item.Short)
	}

	h.Perform(ctx, list, userID)
	return &pb.DeleteUrlsResponse{}, nil
}
