package ws

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Handler struct {
	service *Service
	log     *slog.Logger
}

func NewHandler(stores Stores, log *slog.Logger) *Handler {
	return &Handler{
		service: NewService(stores, log),
		log:     log,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.log.Warn("Upgrade error", slog.String("error", err.Error()))
		return
	}

	clientID := r.URL.Query().Get("id")
	if clientID == "" {
		clientID = generateID()
	}

	client := &Client{
		ID:   clientID,
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  h.service.Hub(),
	}
	client.Hub.Register <- client

	go h.service.WritePump(client)
	go h.service.ReadPump(client)
}

func generateID() string {
	return fmt.Sprintf("client_%d", time.Now().UnixNano())
}
