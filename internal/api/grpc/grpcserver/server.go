package grpcserver

import (
	"fmt"
	"log"
)

func (h *GRPCHandler) Start() {
	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := h.server.Serve(h.listener); err != nil {
		log.Fatal(err)
	}
}

func (h *GRPCHandler) Stop() {
	h.server.Stop()
}
