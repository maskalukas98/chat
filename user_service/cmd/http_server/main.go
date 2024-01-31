package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
	_ "user_service/cmd/http_server/docs"
	"user_service/src/infrastructure/rest"
)

var router = gin.Default()

func main() {

	apiGroup := router.Group("/api")
	v1Group := apiGroup.Group("/v1")

	rest.AddUserRoutes(v1Group)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + os.Getenv("HTTP_PORT"))
}
