package port_output

import "chat_service/src/domain/aggregate"

type MessageRepository interface {
	AddMessage(message aggregate.Message) error
	HasConversationStarted(senderId int64, receiverID int64) bool
	ListRecentMessages(senderId int64, receiverId int64, limit int64) []aggregate.Message
}
