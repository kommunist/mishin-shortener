package interceptors

import (
	"context"
	"mishin-shortener/internal/secure"

	pb "mishin-shortener/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthCheck(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod != pb.DeleteUrls_DeleteUrls_FullMethodName && info.FullMethod != pb.UserUrls_UserUrls_FullMethodName {
		return handler(ctx, req)
	}

	var authToken string

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		value := md.Get("Token")
		if len(value) > 0 {
			authToken = value[0]
		}
	}

	if authToken != "" {
		userID, err := secure.Decrypt(authToken)
		if err != nil || userID == "" { // и если не удалось расшифровать
			return nil, status.Error(codes.Unauthenticated, "check auth failed")
		} else { // единственный положительный сценарий
			newCtx := context.WithValue(ctx, secure.UserIDKey, userID)
			return handler(newCtx, req)
		}
	} else {
		return nil, status.Error(codes.Unauthenticated, "check auth failed")
	}

}
