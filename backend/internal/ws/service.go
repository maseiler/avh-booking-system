package ws

import (
	"context"
	"log/slog"

	"github.com/gorilla/websocket"
)

type Service struct {
	log       *slog.Logger
	hub       *Hub
	validator *messageValidator
	stores    Stores
}

func NewService(stores Stores, log *slog.Logger) *Service {
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
		validator: newMessageValidator(),
		stores:    stores,
	}
}

func (s *Service) Hub() *Hub {
	return s.hub
}

func (s *Service) NewClient(ctx context.Context, cancel context.CancelFunc, conn *websocket.Conn) *Client {
	return &Client{
		Hub:    s.hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Ctx:    ctx,
		Cancel: cancel,
	}
}
