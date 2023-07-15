package main

import (
    "log"
    "net"

    "session-service/pkg/config"
    "session-service/pkg/pb"
    "session-service/pkg/services"
    "google.golang.org/grpc"
)

func main() {
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }


    lis, err := net.Listen("tcp", c.Port)

    if err != nil {
        log.Fatalln("Failed to listing:", err)
    }

    grpcServer := grpc.NewServer()

    server := services.Server{}
    pb.RegisterSessionServiceServer(grpcServer, &server)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalln("Failed to serve:", err)
    }
} 