package value_object

import "strconv"

type ConversationId struct {
	value string
}

func NewConversationId(senderId int64, receiverId int64) *ConversationId {
	return &ConversationId{
		value: create(senderId, receiverId),
	}
}

func (r *ConversationId) GetValue() string {
	return r.value
}

func create(senderId int64, receiverId int64) string {
	smallerId := min(senderId, receiverId)
	largerId := max(senderId, receiverId)

	return strconv.FormatInt(smallerId, 10) + ":" + strconv.FormatInt(largerId, 10)
}
