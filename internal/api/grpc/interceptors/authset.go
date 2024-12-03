package interceptors

import (
	"context"
	"mishin-shortener/internal/secure"

	pb "mishin-shortener/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthSet(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod != pb.Create_Create_FullMethodName && info.FullMethod != pb.CreateBatch_CreateBatch_FullMethodName {
		return handler(ctx, req)
	}

	var userID string
	var authToken string

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		value := md.Get("Token")
		if len(value) > 0 {
			authToken = value[0]
		}
	}

	if authToken != "" {
		var err error

		userID, err = secure.Decrypt(authToken)
		if err != nil || userID == "" {
			userID = newuserID()
		}
	} else {
		userID = newuserID()
	}
	newCtx := context.WithValue(ctx, secure.UserIDKey, userID)

	return handler(newCtx, req) // запускаем работу

}

func newuserID() string {
	return uuid.New().String()
}
