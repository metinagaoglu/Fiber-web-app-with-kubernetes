package auth

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/config"
	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/social/pb"
)

type ServiceClient struct {
	Client pb.SocialServiceClient
}

func InitServiceClient(c *config.Config) pb.SocialServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewSocialServiceClient(cc)
}
