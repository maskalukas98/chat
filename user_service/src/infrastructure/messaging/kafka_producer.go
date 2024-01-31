package messaging

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

type Msg struct {
	Action    string    `json:"action"`
	IpAddress string    `json:"ip_address"`
	Timestamp time.Time `json:"timestamp"`
}

const (
	userEventsTopic = "user_events"
)

func NewKafkaProducer(producer sarama.SyncProducer) *KafkaProducer {
	return &KafkaProducer{
		producer: producer,
	}
}

func (r *KafkaProducer) SendAnalyticUserEvent(action string, userIpAddress string) {
	obj := Msg{
		Action:    action,
		IpAddress: userIpAddress,
		Timestamp: time.Now(),
	}

	bytes, err2 := json.Marshal(obj)

	if err2 != nil {
		log.Fatalf("das")
	}

	message := &sarama.ProducerMessage{
		Topic: userEventsTopic,
		Value: sarama.StringEncoder(bytes),
	}

	_, _, err := r.producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
