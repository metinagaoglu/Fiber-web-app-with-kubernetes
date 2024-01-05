package main

import (
    "log"
    "net"

    "session-service/pkg/config"
    "session-service/pkg/pb"
    "session-service/pkg/services"
    "google.golang.org/grpc"

    db "session-service/pkg/db"
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
    
    rdb := db.ConnectRedis()
    grpcServer := grpc.NewServer()

    server := services.Server{
        Rdb: rdb,
    }
    pb.RegisterSessionServiceServer(grpcServer, &server)

    err = grpcServer.Serve(lis)
    if err != nil {
        log.Fatalln("Failed to serve:", err)
    }

    log.Println("Serving on", c.Port)
} 