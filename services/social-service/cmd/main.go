package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"social-service/pkg/config"
	"social-service/pkg/db"
	"social-service/pkg/pb"
	"social-service/pkg/services"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	h := db.Init(c.Couchbase_Connection, c.Couchbase_Username, c.Couchbase_Password)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSocialServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
