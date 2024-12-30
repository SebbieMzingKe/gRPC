package main

import "log"

func main() {
	httpServer := NewHttpServer(":8001")
	go func() {
        if err := httpServer.Run(); err != nil {
            log.Fatal("HTTP server error:", err)
        }
    }()

	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()
}