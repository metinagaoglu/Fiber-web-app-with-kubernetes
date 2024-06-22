package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/config"
	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/social/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	ginRoutes := r.Group("/social")
	ginRoutes.POST("/add-follower", svc.AddFollower)
	ginRoutes.POST("/remove-follower", svc.RemoveFollower)
	ginRoutes.POST("/get-followers", svc.GetFollowers)
	ginRoutes.POST("/count-followers", svc.CountFollowers)

	return svc
}

func (svc *ServiceClient) AddFollower(ctx *gin.Context) {
	routes.AddFollower(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveFollower(ctx *gin.Context) {
	routes.RemoveFollower(ctx, svc.Client)
}

func (svc *ServiceClient) GetFollowers(ctx *gin.Context) {
	routes.GetFollowers(ctx, svc.Client)
}

func (svc *ServiceClient) CountFollowers(ctx *gin.Context) {
	routes.CountFollowers(ctx, svc.Client)
}
