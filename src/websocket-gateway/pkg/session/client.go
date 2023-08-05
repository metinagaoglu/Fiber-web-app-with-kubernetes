package session

import (
    "fmt"

    config "websocket-gateway/pkg/config"
    "google.golang.org/grpc"
		pb "websocket-gateway/pkg/session/pb"
)

type ServiceClient struct {
  Client pb.SessionServiceClient
}

func InitServiceClient() pb.SessionServiceClient {
    // using WithInsecure() because no SSL running
    //cc, err := grpc.Dial(getSessionServiceUrl(c.AuthSvcUrl), grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
    c, _ := config.LoadConfig()
		cc, err := grpc.Dial(c.SessionSvcUrl,
			grpc.WithInsecure(),
		)

    if err != nil {
        fmt.Println("Could not connect:", err)
    }


    return pb.NewSessionServiceClient(cc)
}