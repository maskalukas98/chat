package mock

import "github.com/stretchr/testify/mock"

type MessagingProducer struct {
	mock.Mock
}

func (m *MessagingProducer) SendAnalyticUserEvent(action string, userIpAddress string) {
	m.Called(action, userIpAddress)
}
