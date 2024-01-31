package repository

import (
	"chat_service/src/domain/aggregate"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisRecentMessagesRepository struct {
	redisClient *redis.Client
}

func NewRedisRecentMessagesRepository(redisClient *redis.Client) *RedisRecentMessagesRepository {
	return &RedisRecentMessagesRepository{
		redisClient: redisClient,
	}
}

type messageDocument struct {
	SenderId   int64     `json:"sender_id"`
	ReceiverId int64     `json:"receiver_id"`
	Message    string    `json:"message"`
	SentAt     time.Time `json:"sent_at"`
}

func (r *RedisRecentMessagesRepository) AppendNewMessage(conversationId string, message aggregate.Message) error {
	message.GetConversationId()
	jsonValue, err := json.Marshal(message)

	if err != nil {
		return err
	}

	r.redisClient.RPush(r.redisClient.Context(), "messages:"+conversationId, jsonValue).Result()

	return nil
}

func (r *RedisRecentMessagesRepository) LoadNextMessages(conversationId string, offset int64, limit int64) []aggregate.Message {
	end := offset + limit - 1

	result, _ := r.redisClient.LRange(r.redisClient.Context(), "messages:"+conversationId, offset, end).Result()

	var messages []aggregate.Message
	for _, jsonStr := range result {
		var msg aggregate.Message
		err := json.Unmarshal([]byte(jsonStr), &msg)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			continue
		}
		messages = append(messages, msg)
	}

	return messages
}

func (r *RedisRecentMessagesRepository) SetMessages(messages []aggregate.Message) {
	if len(messages) == 0 {
		return
	}

	conversationId := messages[0].GetConversationId()
	r.redisClient.Del(context.TODO(), "messages:"+conversationId.GetValue()).Result()

	for _, message := range messages {
		doc := messageDocument{
			SenderId:   message.SenderId,
			ReceiverId: message.ReceiverId,
			Message:    message.Message,
			SentAt:     message.SentAt,
		}

		jsonContent, _ := json.Marshal(doc)
		r.redisClient.LPush(r.redisClient.Context(), "messages:"+conversationId.GetValue(), jsonContent).Result()
	}
}
