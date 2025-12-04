package services

import (
	"encoding/json"
	"log"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	hub *models.Hub
}

type Message struct {
	Type    string `json:"type"`
	Payload interface{}
}
type Query struct {
	Table string `json:"table"`
	Id    int    `json:"id"`
}

func getQueryFromMessage(data interface{}) Query {
	m := data.(map[string]interface{})
	query := Query{}
	if table, ok := m["table"].(string); ok {
		query.Table = table
	}
	return query
}

func NewWebSocketService() *WebSocketService {
	hub := &models.Hub{
		Clients:    make(map[*models.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
	}

	go hub.Run()

	return &WebSocketService{hub: hub}
}

func (ws *WebSocketService) GetHub() *models.Hub {
	return ws.hub
}

func (ws *WebSocketService) BroadcastMessage(message []byte) {
	ws.hub.Broadcast <- message
}

// ReadPump pumps messages from the websocket connection to the hub
func ReadPump(c *models.Client) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512 * 1024) // 512KB max message size

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		// Process message
		//var msg map[string]interface{}
		msg := Message{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("JSON parse error: %v", err)
			continue
		}

		// Handle different message types
		switch msg.Type {
		case "broadcast":
			c.Hub.Broadcast <- message
		case "ping":
			pong, _ := json.Marshal(map[string]string{"type": "pong"})
			c.Send <- pong
		case "query":
			query := getQueryFromMessage(msg.Payload)
			log.Printf("%v", query.Table)
			/*
				switch payload["table"] {
				case "accounts":
					log.Println("Query accounts")
				default:
					panic("Invalid table")
				}
			*/
		default:
			// Echo to sender
			log.Printf("Message: %v", msg)
			c.Send <- message
		}
	}
}

// WritePump pumps messages from the hub to the websocket connection
func WritePump(c *models.Client) {
	defer c.Conn.Close()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// Hub closed the channel
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
