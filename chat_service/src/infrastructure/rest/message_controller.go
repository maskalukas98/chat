package rest

import (
	"chat_service/src/application/command"
	custom_error "chat_service/src/application/error"
	port_input "chat_service/src/domain/port/input"
	"chat_service/src/infrastructure/rest/dto"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MessageController struct {
	sendMessageUseCase port_input.SendMessage
	getMessagesUseCase port_input.GetMessages
}

func NewMessageController(
	sendMessageUseCase *port_input.SendMessage,
	getMessagesUseCase *port_input.GetMessages,
) *MessageController {
	return &MessageController{
		sendMessageUseCase: *sendMessageUseCase,
		getMessagesUseCase: *getMessagesUseCase,
	}
}

func (r *MessageController) GetMessages(c *gin.Context) {
	offsetStr := c.Query("offset")

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil && offsetStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	senderIdStr := c.Query("sender_id")
	senderId, err := strconv.ParseInt(senderIdStr, 10, 64)
	if err != nil && senderIdStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender id"})
		return
	}

	receiverIdStr := c.Query("receiver_id")
	receiverId, err := strconv.ParseInt(receiverIdStr, 10, 64)
	if err != nil && receiverIdStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver id"})
		return
	}

	request := command.GetMessagesRequest{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Offset:     offset,
	}

	response, err := r.getMessagesUseCase.Get(request)

	if err != nil {
		var userNotFoundErr *custom_error.UserNotFoundError
		if errors.As(err, &userNotFoundErr) {
			c.JSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("User not exists with id: %d", request.ReceiverId),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error"))
		return
	}

	responseDto := dto.GetMessagesResponseDto{}

	c.JSON(http.StatusOK, responseDto.CreateFromGetMessagesResponse(*response))
}
