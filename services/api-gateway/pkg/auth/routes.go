package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	ginRoutes := r.Group("/auth")
	ginRoutes.POST("/register", svc.Register)
	ginRoutes.POST("/login", svc.Login)
	ginRoutes.POST("/validate", svc.Validate)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Validate(ctx *gin.Context) {
	routes.Validate(ctx, svc.Client)
}
