package main

import (
	handlers "kitchen/services/orders/handler/orders"
	"kitchen/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()

	// register grpc services
	orderService := service.NewOrderService()
	handlers.NewGrpcOrdersService(gRPCServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return gRPCServer.Serve(lis)
}