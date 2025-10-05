package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

type WSMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

var hub = &Hub{
	clients: make(map[*websocket.Conn]bool),
}

// Tambah client baru
func (h *Hub) AddClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[conn] = true
}

// Hapus client
func (h *Hub) RemoveClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, conn)
	conn.Close()
}

// Broadcast pesan ke semua client
func (h *Hub) Broadcast(msgType string, payload interface{}) {
	message := WSMessage{
		Type:    msgType,
		Payload: payload,
	}
	data, err := json.Marshal(message)
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	for conn := range h.clients {
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Write error:", err)
			h.RemoveClient(conn)
		}
	}
}

// Instance global
func GetHub() *Hub {
	return hub
}
