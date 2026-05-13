package helpers

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WSHub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan interface{}
	Mu        sync.Mutex
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewHub() *WSHub {
	return &WSHub{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan interface{}),
	}
}

func (h *WSHub) Send(data interface{}) {
	h.Broadcast <- data
}

func (h *WSHub) Run() {
	for {
		msg := <-h.Broadcast
		h.Mu.Lock()
		for client := range h.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(h.Clients, client)
			}
		}
		h.Mu.Unlock()
	}
}
