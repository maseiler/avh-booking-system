package ws

import (
	"github.com/av-huette/avh-booking-system/internal/validation"
	"github.com/gorilla/websocket"
	"log/slog"
)

type Service struct {
	log       *slog.Logger
	hub       *Hub
	validator *validation.WebSocketValidator
	stores    Stores
}

func NewService(stores Stores, log *slog.Logger, schemasPath string) *Service {
	hub := &Hub{
		Log:        log,
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}

	go hub.Run()

	return &Service{
		log:       log,
		hub:       hub,
		validator: validation.NewWebSocketValidator(schemasPath),
		stores:    stores,
	}
}

func (s *Service) Hub() *Hub {
	return s.hub
}

func (s *Service) NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Hub:  s.hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
}
