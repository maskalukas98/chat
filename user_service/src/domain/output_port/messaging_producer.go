package output_port

type MessagingProducer interface {
	SendAnalyticUserEvent(action string, userIpAddress string)
}
