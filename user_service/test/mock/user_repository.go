package mock

import (
	"github.com/stretchr/testify/mock"
	"user_service/src/domain"
	"user_service/src/domain/value_object"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) CreateUser(name, email string) (*value_object.UserId, error) {
	args := m.Called(name, email)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*value_object.UserId), args.Error(1)
}

func (m *UserRepository) GetUserById(userId value_object.UserId) (*domain.User, error) {
	args := m.Called(&userId)

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	return args.Get(0).(*domain.User), args.Error(1)
}
