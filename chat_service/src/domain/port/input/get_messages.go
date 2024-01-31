package port_input

import "chat_service/src/application/command"

type GetMessages interface {
	Get(request command.GetMessagesRequest) (*command.GetMessagesResponse, error)
}
