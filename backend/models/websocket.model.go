package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSHub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan interface{}
	Mu        sync.Mutex
	Name      string
}

type WebsocketIsland struct {
	CPUHub    *WSHub
	LoggerHub *WSHub
}
