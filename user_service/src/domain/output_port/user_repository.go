package output_port

import (
	"user_service/src/domain"
	"user_service/src/domain/value_object"
)

type UserRepositoryPort interface {
	CreateUser(username, email string) (*value_object.UserId, error)
	GetUserById(userId value_object.UserId) (*domain.User, error)
}
