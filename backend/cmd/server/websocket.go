package main

import (
	"encoding/json"
	"log"

	"github.com/go-playground/validator/v10"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	hub *models.Hub
}

func getQueryFromMessage(data interface{}) (*models.Query, *models.WsError) {
	// Convert the interface{} back to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal interface"}
	}

	// Unmarshal directly into the Query struct
	var query models.Query
	if err := json.Unmarshal(jsonBytes, &query); err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal to Query"}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(query)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadStruct, Message: err.Error(), Details: "Invalid struct"}
	}

	return &query, nil
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

// processQuery unmarshals the message, fetches the data from the database and returns the object as JSON
func processQuery(message models.Message, dbModels *dbModels) ([]byte, *models.WsError) {
	query, err := getQueryFromMessage(message.Payload)
	if err != nil {
		return nil, err
	}

	switch query.Table {
	case "account":
		{
			accounts, _ := dbModels.account.Get(query)
			b, err := json.Marshal(map[string]interface{}{
				"type":     "accounts",
				"accounts": accounts,
			})
			if err != nil {
				return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal []Account"}
			}
			return b, nil
		}
	case "product":
		// ...
	}

	return nil, &models.WsError{Code: models.WsUnknown, Message: "Unknown error", Details: "Could not process query"}
}

// ReadPump pumps messages from the websocket connection to the hub
func ReadPump(c *models.Client, dbModels *dbModels) {
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
		msg := models.Message{}
		if err := json.Unmarshal(message, &msg); err != nil {
			wsErr := models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal message"}
			log.Printf("%v", wsErr)
			errBytes, _ := json.Marshal(wsErr)
			c.Send <- errBytes
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
			b, err := processQuery(msg, dbModels)
			if err != nil {
				errBytes, _ := json.Marshal(err)
				c.Send <- errBytes
				continue
			}
			c.Send <- b
		case "mutation":
			/*
				b, err := processMutation(msg, dbModels)
				if err != nil {
					log.Printf("Mutation error: %v", err)
					// TODO send error to client
					errBytes, _ := json.Marshal(err)
					c.Send <- errBytes
					continue
				}
				c.Send <- true
			*/
		default:
			// Echo to sender
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
