package rest

import (
	"chat_service/src"
	"github.com/gin-gonic/gin"
)

func AddMessagesRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/messages")

	messageController := NewMessageController(&src.SendMessageUseCase, &src.GetMessagesUseCase)

	userGroup.GET("", messageController.GetMessages)
}
