package main

import (
	"chat_service/src/infrastructure/rest"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var router = gin.Default()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1")

	rest.AddMessagesRoutes(v1)
	router.Run(":" + os.Getenv("HTTP_PORT"))
}
