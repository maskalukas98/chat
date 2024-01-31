package messaging

import (
	"fmt"
	"github.com/IBM/sarama"
)

type KafkaProducerFactory struct {
}

func (r *KafkaProducerFactory) CreateSyncProducer(brokerAddress string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, config)

	if err != nil {
		panic(fmt.Sprintf("cannot create kafka producer: %v", err))
	}

	return producer
}
