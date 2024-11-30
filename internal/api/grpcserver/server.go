package grpcserver

import (
	"fmt"
	"log"
)

func (h *GRPCHandler) Start() {
	// listener, err := net.Listen("tcp", ":3200")
	// if err != nil {
	// 	slog.Error("Error when listen net", "err", err)
	// }
	// // h.listener = &listener

	// // создаём gRPC-сервер без зарегистрированной службы
	// s := grpc.NewServer()
	// // регистрируем сервис
	// pb.RegisterShortenerServer(s, h)

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := h.server.Serve(h.listener); err != nil {
		log.Fatal(err)
	}
}

func (h *GRPCHandler) Stop() {
	// err := listener.Close()
	// slog.Error("Error when close listner", "err", err)
}
