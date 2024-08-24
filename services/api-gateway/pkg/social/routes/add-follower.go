package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/social/pb"
)

type AddFollowerRequestBody struct {
	FolloweeId int `json:"followee_id" binding:"required"`
}

func AddFollower(ctx *gin.Context, c pb.SocialServiceClient) {
	body := AddFollowerRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddFollower(ctx, &pb.AddFollowerRequest{
		Followee: int64(body.FolloweeId),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
