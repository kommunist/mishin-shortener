package deleteurls

import (
	"context"
	pb "mishin-shortener/proto"
)

const defaultUserID = "ququ"

func (h *Handler) CallGRPC(ctx context.Context, in *pb.DeleteUrlsRequest) (*pb.DeleteUrlsResponse, error) {
	list := make([]string, 0, len(in.List))
	for _, item := range in.List {
		list = append(list, item.Short)
	}

	h.Perform(ctx, list, defaultUserID)
	return &pb.DeleteUrlsResponse{}, nil
}
