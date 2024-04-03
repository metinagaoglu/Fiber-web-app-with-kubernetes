package auth

import (
	"fmt"
	"math/rand"
	"strings"

	"google.golang.org/grpc"
	"websocket-gateway/pkg/auth/pb"
	config "websocket-gateway/pkg/config"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	//cc, err := grpc.Dial(getAuthServiceUrl(c.AuthSvcUrl), grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
	c, _ := config.LoadConfig()
	cc, err := grpc.Dial(getAuthServiceUrl(c.AuthSvcUrl),
		grpc.WithInsecure(),
	)

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}

func getAuthServiceUrl(url string) string {
	// split url by , and get random element
	urls := strings.Split(url, ",")
	url = urls[rand.Intn(len(urls))]

	// return url
	return url
}
