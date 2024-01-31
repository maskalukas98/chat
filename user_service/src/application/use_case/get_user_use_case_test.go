package use_case

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user_service/src/application/command"
	"user_service/src/domain"
	"user_service/src/domain/output_port"
	"user_service/src/domain/value_object"
	custom_mock "user_service/test/mock"
)

func TestGetUserUseCase_GetUserById_Successful(t *testing.T) {
	// prepare
	request := command.GetUserByIdRequest{
		UserId: int64(40),
	}

	var userRepository output_port.UserRepositoryPort = new(custom_mock.UserRepository)
	userRepository.(*custom_mock.UserRepository).
		On("GetUserById", value_object.NewUserId(request.UserId)).
		Return(&domain.User{
			UserId: *value_object.NewUserId(request.UserId),
			Name:   "lukas",
			Email:  "lukas.maska@fakeEmail.com",
		}, nil)

	// execute
	getUserUseCase := NewGetUserUseCase(&userRepository)
	response, err := getUserUseCase.GetById(request)

	// validate
	assert.Equal(t, response.UserId, int64(40))
	assert.Equal(t, response.Name, "lukas")
	assert.Equal(t, response.Email, "lukas.maska@fakeEmail.com")
	assert.Nil(t, err)
}

func TestGetUserUseCase_GetUserById_NotFound(t *testing.T) {
	// prepare
	request := command.GetUserByIdRequest{
		UserId: int64(40),
	}

	var userRepository output_port.UserRepositoryPort = new(custom_mock.UserRepository)
	userRepository.(*custom_mock.UserRepository).
		On("GetUserById", value_object.NewUserId(request.UserId)).
		Return(nil, nil)

	// execute
	getUserUseCase := NewGetUserUseCase(&userRepository)
	response, err := getUserUseCase.GetById(request)

	// validate
	assert.Nil(t, response)
	assert.Nil(t, err)
}
