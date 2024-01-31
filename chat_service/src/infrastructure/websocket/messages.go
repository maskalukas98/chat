package infrastructure_websocket

import (
	"chat_service/src"
	"chat_service/src/application/command"
	"chat_service/src/domain/aggregate"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
)

type MessageHandler struct {
}

var upgrader = websocket.Upgrader{}

type Client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

var clients = make(map[int64]*Client)
var broadcast = make(chan aggregate.Message)

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	senderID, _ := strconv.ParseInt(r.Header.Get("client_id"), 10, 64)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &Client{conn: conn}

	client.mu.Lock()
	clients[senderID] = client
	client.mu.Unlock()

	fmt.Println("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			client.mu.Lock()
			delete(clients, 2)
			client.mu.Unlock()
			break
		}

		fmt.Println(messageType)
		fmt.Printf("Received message: %s\n", p)

		var receivedMessage aggregate.Message
		err = json.Unmarshal(p, &receivedMessage)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}

		src.SendMessageUseCase.Send(command.SendMessageRequest{
			ReceiverId: receivedMessage.ReceiverId,
			SenderId:   senderID,
			Message:    receivedMessage.Message,
		})

		broadcast <- receivedMessage
	}
}

func HandleMessages() {
	for {
		message := <-broadcast
		client := clients[message.ReceiverId]

		if clients[message.ReceiverId] == nil {
			return
		}

		byteSlice, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		client.mu.Lock()
		err = client.conn.WriteMessage(websocket.TextMessage, byteSlice)
		client.mu.Unlock()
	}
}
