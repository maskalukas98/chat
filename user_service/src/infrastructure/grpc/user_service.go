package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"user_service/src/application/command"
	custom_error "user_service/src/application/error"
	"user_service/src/domain/input_port"
	"user_service/src/infrastructure/grpc/gen"
)

type UserService struct {
	createUserUseCase input_port.CreateUserPort
	getUserUseCase    input_port.GetUserPort
}

func NewUserService(
	createUserUseCase *input_port.CreateUserPort,
	getUserUseCase *input_port.GetUserPort,
) *UserService {
	return &UserService{
		createUserUseCase: *createUserUseCase,
		getUserUseCase:    *getUserUseCase,
	}
}

func (r *UserService) CreateUser(ctx context.Context, message *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	peer, ok := peer.FromContext(ctx)

	if ok == false {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	var request = command.CreateUserRequest{
		Name:      message.Name,
		Email:     message.Email,
		IpAddress: peer.Addr.String(),
	}

	response, err := r.createUserUseCase.Create(request)

	if err != nil {
		var userExistsErr *custom_error.UserAlreadyExistsError
		if errors.As(err, &userExistsErr) {
			return nil, status.Errorf(codes.AlreadyExists, "User already exists with email: %s", message.Email)
		}

		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &gen.CreateUserResponse{
		UserId: response.UserId,
	}, nil
}

func (r *UserService) GetUserById(_ context.Context, message *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	request := command.GetUserByIdRequest{UserId: message.GetUserId()}

	response, err := r.getUserUseCase.GetById(request)

	if response == nil && err == nil {
		return nil, status.Errorf(codes.NotFound, "User with ID %d not found.", message.GetUserId())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &gen.GetUserResponse{
		UserId: response.UserId,
		Name:   response.Name,
		Email:  response.Email,
	}, nil
}
