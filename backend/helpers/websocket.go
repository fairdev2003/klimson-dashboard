package helpers

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/zgierz/klimson/backend/logger"
)

type WSHub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan interface{}
	Mu        sync.Mutex
	Name      string
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		logger.GreenServerLog("Origin: ", "s")
		return true
	},
}

func NewHub(name string) *WSHub {
	return &WSHub{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan interface{}),
		Name:      name,
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
