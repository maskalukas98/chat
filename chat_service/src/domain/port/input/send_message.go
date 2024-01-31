package port_input

import "chat_service/src/application/command"

type SendMessage interface {
	Send(request command.SendMessageRequest) error
}
