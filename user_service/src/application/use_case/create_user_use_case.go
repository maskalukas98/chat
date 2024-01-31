package use_case

import (
	"user_service/src/application/command"
	_const "user_service/src/domain/const"
	"user_service/src/domain/output_port"
)

type CreateUserUseCase struct {
	userRepository    output_port.UserRepositoryPort
	messagingProducer output_port.MessagingProducer
}

func NewCreateUserUseCase(
	userRepository *output_port.UserRepositoryPort,
	messagingProducer *output_port.MessagingProducer,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository:    *userRepository,
		messagingProducer: *messagingProducer,
	}
}

func (r *CreateUserUseCase) Create(request command.CreateUserRequest) (*command.CreateUserResponse, error) {
	userId, err := r.userRepository.CreateUser(request.Name, request.Email)

	if err != nil {
		return nil, err
	}

	r.messagingProducer.SendAnalyticUserEvent(_const.UserAction.Created, request.IpAddress)

	return &command.CreateUserResponse{
		UserId: userId.GetId(),
	}, nil
}
