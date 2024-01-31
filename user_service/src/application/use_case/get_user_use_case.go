package use_case

import (
	"user_service/src/application/command"
	"user_service/src/domain/output_port"
	"user_service/src/domain/value_object"
)

type GetUserUseCase struct {
	userRepository output_port.UserRepositoryPort
}

func NewGetUserUseCase(userRepository *output_port.UserRepositoryPort) *GetUserUseCase {
	return &GetUserUseCase{
		userRepository: *userRepository,
	}
}

func (r *GetUserUseCase) GetById(request command.GetUserByIdRequest) (*command.GetUserResponse, error) {
	user, err := r.userRepository.GetUserById(*value_object.NewUserId(request.UserId))

	if user == nil && err == nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &command.GetUserResponse{
		UserId: user.UserId.GetId(),
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}
