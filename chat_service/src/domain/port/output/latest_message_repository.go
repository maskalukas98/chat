package port_output

import "chat_service/src/domain/aggregate"

type LatestMessagesCacheRepository interface {
	AppendNewMessage(conversationId string, message aggregate.Message) error
	LoadNextMessages(conversationId string, offset int64, limit int64) []aggregate.Message
	SetMessages(messages []aggregate.Message)
}
