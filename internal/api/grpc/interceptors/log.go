package interceptors

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func Log(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req) // запускаем работу

	duration := time.Since(start)

	inf := status.Convert(err) // парсим ошибку

	var code, msg string

	if err != nil {
		code = inf.Code().String()
		msg = inf.Message()
	} else {
		code = "Success"
	}

	slog.Info("Request GRPC", "err", err, "code", code, "msg", msg, "duration", duration)

	return resp, err
}
