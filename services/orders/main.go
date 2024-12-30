package main

import (
	"log"
	"os"
)

// func main() {
// 	httpServer := NewHttpServer(":8001")
// 	go func() {
//         if err := httpServer.Run(); err != nil {
//             log.Fatal("HTTP server error:", err)
//         }
//     }()

// 	grpcServer := NewGRPCServer(":9000")
// 	grpcServer.Run()
// }

func main() {
    httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = "8001"
    }
    grpcPort := os.Getenv("GRPC_PORT")
    if grpcPort == "" {
        grpcPort = "9000"
    }

    httpServer := NewHttpServer(":" + httpPort)
    go func() {
        if err := httpServer.Run(); err != nil {
            log.Fatal("HTTP server error:", err)
        }
    }()

    grpcServer := NewGRPCServer(":" + grpcPort)
    grpcServer.Run()
}