package main

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	return conn
}

// func main() {
// 	httpServer := NewHttpServer(":2000")
// 	if err := httpServer.Run(); err != nil {
// 		log.Fatal("server failed to start:", err)
// 	}
	
// }

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "2000"
    }
    httpServer := NewHttpServer(":" + port)
    if err := httpServer.Run(); err != nil {
        log.Fatal("server failed to start:", err)
    }
}