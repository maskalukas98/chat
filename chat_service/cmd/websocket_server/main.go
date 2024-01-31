package main

import (
	infrastructure_websocket "chat_service/src/infrastructure/websocket"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go infrastructure_websocket.HandleMessages()

	http.HandleFunc("/ws", infrastructure_websocket.HandleConnection)

	err = http.ListenAndServe(":"+os.Getenv("WS_PORT"), nil)
	if err != nil {
		fmt.Println(err)
	}
}
