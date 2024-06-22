package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/auth"
	"github.com/metinagaoglu/go-grpc-api-gateway/pkg/config"
	social "github.com/metinagaoglu/go-grpc-api-gateway/pkg/social"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	//TODO: Allow cors
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth.RegisterRoutes(r, &c)
	social.RegisterRoutes(r, &c)
	r.Run(c.Port)
}
