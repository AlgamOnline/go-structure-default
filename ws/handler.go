package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	GetHub().AddClient(conn)
	log.Println("Client connected")

	defer GetHub().RemoveClient(conn)

	for {
		// jika client kirim sesuatu (optional kita bisa handle)
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected")
			break
		}
	}
}
