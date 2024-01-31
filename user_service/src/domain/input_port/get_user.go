package input_port

import "user_service/src/application/command"

type GetUserPort interface {
	GetById(request command.GetUserByIdRequest) (*command.GetUserResponse, error)
}
