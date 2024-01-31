package mock

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/mock"
)

type SyncProducer struct {
	mock.Mock
}

func (s SyncProducer) SendMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	s.Called(message)
	return 1, 1, nil
}

func (s SyncProducer) SendMessages(msgs []*sarama.ProducerMessage) error {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) TxnStatus() sarama.ProducerTxnStatusFlag {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) IsTransactional() bool {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) BeginTxn() error {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) CommitTxn() error {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) AbortTxn() error {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) AddOffsetsToTxn(offsets map[string][]*sarama.PartitionOffsetMetadata, groupId string) error {
	//TODO implement me
	panic("implement me")
}

func (s SyncProducer) AddMessageToTxn(msg *sarama.ConsumerMessage, groupId string, metadata *string) error {
	//TODO implement me
	panic("implement me")
}
