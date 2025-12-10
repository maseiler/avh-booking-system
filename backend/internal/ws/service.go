package ws

import (
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/validation"
	"github.com/gorilla/websocket"
)

type Service struct {
	hub       *models.Hub
	validator *validation.WebSocketValidator
	dbModels  *models.DbModels
}

func NewService(dbModels *models.DbModels) *Service {
	hub := &models.Hub{
		Clients:    make(map[*models.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
	}

	go hub.Run()

	return &Service{
		hub:       hub,
		validator: validation.NewWebSocketValidator(),
		dbModels:  dbModels,
	}
}

func (s *Service) Hub() *models.Hub {
	return s.hub
}

func (s *Service) NewClient(conn *websocket.Conn) *models.Client {
	return &models.Client{
		Hub:  s.hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
}
