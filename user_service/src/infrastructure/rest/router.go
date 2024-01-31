package rest

import (
	"github.com/gin-gonic/gin"
	"user_service/src"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")

	userController := NewUserController(&src.CreateUserUseCase, &src.GetUserUseCase)

	userGroup.GET(":id", userController.GetUserById)
	userGroup.POST("", userController.CreateUser)
	userGroup.OPTIONS("", userController.GetMethodsForUsers)
	userGroup.OPTIONS(":id", userController.GetMethodsForOneUser)
}
