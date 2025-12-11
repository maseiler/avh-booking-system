package ws

import (
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/validation"
	"github.com/gorilla/websocket"
	"log/slog"
)

type Service struct {
	log       *slog.Logger
	hub       *models.Hub
	validator *validation.WebSocketValidator
	dbModels  *models.DbModels
}

func NewService(dbModels *models.DbModels, log *slog.Logger) *Service {
	hub := &models.Hub{
		Log:        log,
		Clients:    make(map[*models.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
	}

	go hub.Run()

	return &Service{
		log:       log,
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
