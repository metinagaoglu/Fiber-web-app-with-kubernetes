package routes

import (
	"fmt"
	"context"

		auth "websocket-gateway/pkg/auth"
		pb "websocket-gateway/pkg/auth/pb"
)

func AuthByToken(token string) (int64, error) {

	client := auth.InitServiceClient()
	res, err := client.Validate(context.Background(), &pb.ValidateRequest{
			Token: token,
	})

	if err != nil {
			//ctx.AbortWithError(http.StatusBadGateway, err)
			return 0, err
	}
	
	// if res.err != nil {
	// 	fmt.Println("err:", err)
	// 	//ctx.AbortWithError(http.StatusBadGateway, err)
	// 	return false
	// }

	//TODO: add userId

	if res.Status != 200 {
		err := fmt.Errorf("Bad gateway")
		return 0, err
	}

	return res.UserId, nil
}