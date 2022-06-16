package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		Mine("test", "testdata", "21e8", 0, true)
		if err := conn.WriteMessage(messageType, []byte("W")); err != nil {
			log.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	log.Println("Client connected")

	msg := []byte("Talking to Client")
	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println(err)
	}
	// Hashwall(w, r)
	sendMessage(ws)
	reader(ws)
}

func sendMessage(conn *websocket.Conn) {
	msg := []byte("Hello from the server!")
	err := conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println(err)
	}
}
