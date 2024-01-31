package grpc

import (
	"google.golang.org/grpc"
	"user_service/src"
	"user_service/src/infrastructure/grpc/gen"
)

func RegisterGrpcServices(server *grpc.Server) {
	userService := NewUserService(&src.CreateUserUseCase, &src.GetUserUseCase)

	gen.RegisterUserServiceServer(server, userService)
}
