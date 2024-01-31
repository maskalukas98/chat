package input_port

import (
	"user_service/src/application/command"
)

type CreateUserPort interface {
	Create(request command.CreateUserRequest) (*command.CreateUserResponse, error)
}
