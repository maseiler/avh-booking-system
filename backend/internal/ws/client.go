package ws

import (
	"context"

	"github.com/gorilla/websocket"
)

// Client represents a WebSocket client
type Client struct {
	ID     string
	Conn   *websocket.Conn
	Send   chan []byte
	Hub    *Hub
	Ctx    context.Context
	Cancel context.CancelFunc
}
