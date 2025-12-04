package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/services"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (app *application) getAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	account, err := app.dbModels.account.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		app.log.Error("Could not marshal account")

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		app.log.Error("Could not write response")
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Configure origin checking for production
		return true
	},
}

type WebSocketHandler struct {
	wsService *services.WebSocketService
}

func NewWebSocketHandler(wsService *services.WebSocketService) *WebSocketHandler {
	return &WebSocketHandler{wsService: wsService}
}

func (h *WebSocketHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}

	// Get client ID from query params or generate
	clientID := r.URL.Query().Get("id")
	if clientID == "" {
		clientID = generateID()
	}

	client := &models.Client{
		ID:   clientID,
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  h.wsService.GetHub(),
	}

	client.Hub.Register <- client

	// Start goroutines for reading and writing
	go services.WritePump(client)
	go services.ReadPump(client)
}

func generateID() string {
	// Generate unique ID (simplified)
	return fmt.Sprintf("client_%d", time.Now().UnixNano())
}
