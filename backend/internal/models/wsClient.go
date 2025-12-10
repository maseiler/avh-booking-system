package models

import (
	"github.com/av-huette/avh-booking-system/internal/validation"
	"github.com/gorilla/websocket"
)

// Client represents a WebSocket client
type Client struct {
	ID        string
	Conn      *websocket.Conn
	Send      chan []byte
	Hub       *Hub
	Validator *validation.WebSocketValidator
}
