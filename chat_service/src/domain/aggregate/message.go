package aggregate

import (
	"chat_service/src/domain/value_object"
	"time"
)

type Message struct {
	SenderId   int64     `json:"sender_id"`
	ReceiverId int64     `json:"receiver_id"`
	Message    string    `json:"message"`
	SentAt     time.Time `json:"sent_at"`
}

func (r *Message) GetConversationId() value_object.ConversationId {
	return *value_object.NewConversationId(r.SenderId, r.ReceiverId)
}
