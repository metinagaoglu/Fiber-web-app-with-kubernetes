package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/auth/pb"
)

type ValidateRequestBody struct {
	Token string `json:"token" binding:"required"`
}

func Validate(ctx *gin.Context, c pb.AuthServiceClient) {
	body := ValidateRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Validate(context.Background(), &pb.ValidateRequest{
		Token: body.Token,
	})

	if err != nil {
		fmt.Println("=====================")
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
