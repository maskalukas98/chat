package use_case

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"user_service/src/application/command"
	custom_error "user_service/src/application/error"
	"user_service/src/domain/output_port"
	"user_service/src/domain/value_object"
	custom_mock "user_service/test/mock"
)

var messagingProducer output_port.MessagingProducer

func setupTest() {
	messagingProducer = new(custom_mock.MessagingProducer)
	messagingProducer.(*custom_mock.MessagingProducer).
		On("SendAnalyticUserEvent", mock.Anything, mock.Anything).
		Return()
}

func TestCreateUserUseCase_Create_SuccessfulCreation(t *testing.T) {
	// prepare
	setupTest()

	request := command.CreateUserRequest{
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		IpAddress: "192.168.0.1",
	}

	var userRepository output_port.UserRepositoryPort = new(custom_mock.UserRepository)
	userRepository.(*custom_mock.UserRepository).
		On("CreateUser", request.Name, request.Email).
		Return(value_object.NewUserId(20), nil)

	// execute
	createUserUseCase := NewCreateUserUseCase(&userRepository, &messagingProducer)
	response, err := createUserUseCase.Create(request)

	// validate
	assert.Equal(t, response.UserId, int64(20))
	assert.Equal(t, err, nil)
}

func TestCreateUserUseCase_Create_FailureCreation(t *testing.T) {
	// prepare
	setupTest()

	request := command.CreateUserRequest{}

	var userRepository output_port.UserRepositoryPort = new(custom_mock.UserRepository)
	userRepository.(*custom_mock.UserRepository).
		On("CreateUser", mock.Anything, mock.Anything).
		Return(nil, &custom_error.UserAlreadyExistsError{})

	// execute
	createUserUseCase := NewCreateUserUseCase(&userRepository, &messagingProducer)
	response, err := createUserUseCase.Create(request)

	// validate
	assert.Nil(t, response)
	var userExistsErr *custom_error.UserAlreadyExistsError
	assert.True(t, errors.As(err, &userExistsErr))
}
