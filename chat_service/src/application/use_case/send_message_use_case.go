package use_case

import (
	"chat_service/src/application/command"
	custom_error "chat_service/src/application/error"
	"chat_service/src/domain/aggregate"
	port_output "chat_service/src/domain/port/output"
)

type SendMessageUseCase struct {
	messageRepository             port_output.MessageRepository
	recentMessagesCacheRepository port_output.LatestMessagesCacheRepository
	userService                   port_output.UserService
}

func NewSendMessageUseCase(
	messageRepository *port_output.MessageRepository,
	recentMessagesCacheRepository *port_output.LatestMessagesCacheRepository,
	userService *port_output.UserService,

) *SendMessageUseCase {
	return &SendMessageUseCase{
		messageRepository:             *messageRepository,
		recentMessagesCacheRepository: *recentMessagesCacheRepository,
		userService:                   *userService,
	}
}

func (r *SendMessageUseCase) Send(request command.SendMessageRequest) error {
	receiver := r.userService.GetUserById(request.ReceiverId)

	if receiver == nil {
		return &custom_error.UserNotFoundError{UserId: receiver.Id}
	}

	m := &aggregate.Message{
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Message:    request.Message,
	}
	id := m.GetConversationId()

	err := r.messageRepository.AddMessage(*m)

	if err != nil {
		return err
	}

	err = r.recentMessagesCacheRepository.AppendNewMessage(id.GetValue(), *m)

	if err != nil {
		return err
	}

	return nil
}
