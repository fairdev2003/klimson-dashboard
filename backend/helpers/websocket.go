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

func NewHub(name string) *WSHub {
	return &WSHub{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan interface{}),
		Name:      name,
	}
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		logger.GreenServerLog("Origin: ", "s")
		return true
	},
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

type WebsocketRegistry struct {
	mu   sync.RWMutex
	hubs map[string]*WSHub
}

func NewRegistry() *WebsocketRegistry {
	return &WebsocketRegistry{
		hubs: make(map[string]*WSHub),
	}
}

func (r *WebsocketRegistry) Register(path string, hub *WSHub) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.hubs[path] = hub
	go hub.Run()
}

func (r *WebsocketRegistry) GetHub(path string) *WSHub {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.hubs[path]
}

func (r *WebsocketRegistry) InitializeHubs(names ...string) {
	for _, name := range names {
		path := "/ws/" + name
		r.Register(path, NewHub(name))
	}
}
func (r *WebsocketRegistry) InitializeHub(name string) {
	path := "/ws/" + name
	r.Register(path, NewHub(name))

}
