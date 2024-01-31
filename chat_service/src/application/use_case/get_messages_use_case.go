package use_case

import (
	"chat_service/src/application/command"
	custom_error "chat_service/src/application/error"
	port_output "chat_service/src/domain/port/output"
	"chat_service/src/domain/value_object"
)

type GetMessagesUseCase struct {
	messageRepository             port_output.MessageRepository
	recentMessagesCacheRepository port_output.LatestMessagesCacheRepository
	userService                   port_output.UserService
}

func NewGetMessagesUseCase(
	messageRepository *port_output.MessageRepository,
	recentMessagesCacheRepository *port_output.LatestMessagesCacheRepository,
	userService *port_output.UserService,

) *GetMessagesUseCase {
	return &GetMessagesUseCase{
		messageRepository:             *messageRepository,
		recentMessagesCacheRepository: *recentMessagesCacheRepository,
		userService:                   *userService,
	}
}

func (r *GetMessagesUseCase) Get(request command.GetMessagesRequest) (*command.GetMessagesResponse, error) {
	receiver := r.userService.GetUserById(request.ReceiverId)

	if receiver == nil {
		return nil, &custom_error.UserNotFoundError{UserId: request.ReceiverId}
	}

	conversationId := value_object.NewConversationId(request.SenderId, request.ReceiverId)

	messages := r.recentMessagesCacheRepository.LoadNextMessages(conversationId.GetValue(), request.Offset, int64(5))

	response := command.GetMessagesResponse{}
	return response.CreateFromAggregateMessages(messages), nil
}
