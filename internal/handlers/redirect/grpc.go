package redirect

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/errors/deleted"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	to, err := h.Perform(ctx, in.Short)

	if err == nil {
		return &pb.GetResponse{Original: to}, nil
	}

	if _, ok := err.(*deleted.DeletedError); ok { // если удаленнный
		slog.Error("Short url already deleted", "err", err)
		return nil, status.Error(codes.OutOfRange, "Error when call service")
	} else {
		slog.Error("Error when get stats from db", "err", err)
		return nil, status.Error(codes.NotFound, "Error when call service")
	}
}
