package command

type SendMessageRequest struct {
	SenderId   int64
	ReceiverId int64
	Message    string
}
