package command

import "chat_service/src/domain/aggregate"

type GetMessagesRequest struct {
	SenderId   int64
	ReceiverId int64
	Offset     int64
}

type GetMessagesResponse struct {
	Messages []MessageResponse
}

type MessageResponse struct {
	SenderId   int64
	ReceiverId int64
	Message    string
}

func (r *GetMessagesResponse) CreateFromAggregateMessages(messages []aggregate.Message) *GetMessagesResponse {
	newMessages := make([]MessageResponse, 0, len(messages))

	for _, message := range messages {
		newMessages = append(newMessages, MessageResponse{
			SenderId:   message.SenderId,
			ReceiverId: message.ReceiverId,
			Message:    message.Message,
		})
	}

	return &GetMessagesResponse{
		Messages: newMessages,
	}
}
