package routes

import (
	"fmt"
	"context"

		auth "websocket-gateway/pkg/auth"
		pb "websocket-gateway/pkg/auth/pb"
)

func AuthByToken(token string) bool {

		fmt.Println(token)
		if len(token) == 0 {
			return false
		}


	client := auth.InitServiceClient()
	res, err := client.Validate(context.Background(), &pb.ValidateRequest{
			Token:    token,
	})

	if err != nil {
			fmt.Println("err:", err)
			//ctx.AbortWithError(http.StatusBadGateway, err)
			return false
	}
	
	// if res.err != nil {
	// 	fmt.Println("err:", err)
	// 	//ctx.AbortWithError(http.StatusBadGateway, err)
	// 	return false
	// }

	//TODO: add userId

	return res.Status == 200
}