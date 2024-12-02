package simplecreate

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/errors/exist"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CallGRPC(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	var userID string
	if ctx.Value(secure.UserIDKey) == nil {
		return nil, status.Error(codes.Unknown, "Error with auth")
	} else {
		userID = ctx.Value(secure.UserIDKey).(string)
	}

	short, err := h.Perform(ctx, []byte(req.Original), userID)

	if err != nil {
		if _, ok := err.(*exist.ExistError); ok { // обрабатываем проблему, когда уже есть в базе
			return nil, status.Error(codes.AlreadyExists, "Error when call service")

		} else {
			slog.Error("push to storage error", "err", err)
			return nil, status.Error(codes.Unknown, "Error when call service")
		}

	}

	return &pb.CreateResponse{Short: short}, nil
}
