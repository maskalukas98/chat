package messaging

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/mock"
	"testing"
	mock2 "user_service/test/mock"
)

func TestKafkaProducer_SendAnalyticUserEvent_Success(t *testing.T) {
	// prepare
	var syncProducer sarama.SyncProducer = new(mock2.SyncProducer)
	var mockSyncProducer = syncProducer.(*mock2.SyncProducer)

	producer := NewKafkaProducer(syncProducer)

	// checker
	mockSyncProducer.On("SendMessage", mock.MatchedBy(func(arg *sarama.ProducerMessage) bool {
		decodedMessage := decodeMessage(arg)

		if decodedMessage.Action == "created" && decodedMessage.IpAddress == "127.0.0.1:56751" {
			return true
		}

		return false
	}))

	// execute
	producer.SendAnalyticUserEvent("created", "127.0.0.1:56751")

	// validation
	mockSyncProducer.AssertExpectations(t)
}

func decodeMessage(arg *sarama.ProducerMessage) Msg {
	data, _ := arg.Value.Encode()

	var decodedMessage Msg
	json.Unmarshal(data, &decodedMessage)

	return decodedMessage
}
