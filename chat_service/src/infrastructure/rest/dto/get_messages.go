package dto

import (
	"chat_service/src/application/command"
)

type GetMessagesResponseDto struct {
	Messages []MessageResponseDto `json:"messages"`
}

type MessageResponseDto struct {
	SenderId   int64  `json:"sender_id"`
	ReceiverId int64  `json:"receiver_id"`
	Message    string `json:"message"`
}

func (r *GetMessagesResponseDto) CreateFromGetMessagesResponse(response command.GetMessagesResponse) GetMessagesResponseDto {
	newMessages := make([]MessageResponseDto, 0, len(response.Messages))

	for _, message := range response.Messages {
		newMessages = append(newMessages, MessageResponseDto{
			SenderId:   message.SenderId,
			ReceiverId: message.ReceiverId,
			Message:    message.Message,
		})
	}

	return GetMessagesResponseDto{
		Messages: newMessages,
	}
}
