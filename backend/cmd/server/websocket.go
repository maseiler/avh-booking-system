package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	hub *models.Hub
}

func getQueryFromMessage(data interface{}) (*models.Query, error) {
	// Convert the interface{} back to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	// Unmarshal directly into the Query struct
	var query models.Query
	if err := json.Unmarshal(jsonBytes, &query); err != nil {
		return nil, fmt.Errorf("failed to unmarshal query: %w", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	/*
		err = validate.RegisterValidation("validate_operator", func(fl validator.FieldLevel) bool {
			value := fl.Field().Interface().(string)
			switch value {
			case "eq":
				fallthrough
			case "gt":
				fallthrough
			case "gte":
				fallthrough
			case "lt":
				fallthrough
			case "lte":
				fallthrough
			case "ne":
				return true
			default:
				return false
			}
		})
		if err != nil {
			fmt.Printf("Err(s):\n%+v\n", err)
		}
	*/
	err = validate.Struct(query)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
	/*
		for i, filter := range query.Filter {
			err = validate.Struct(filter)
			if err != nil {
				fmt.Printf("Error validating filter %d: %+v\n", i, err)
			} else {
				fmt.Printf("Filter %d ok\n", i)
			}
		}
	*/
	fmt.Printf("XXXXX: %v\n", query)
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
func processQuery(message models.Message, dbModels *dbModels) ([]byte, error) {
	query, err := getQueryFromMessage(message.Payload)
	if err != nil {
		return nil, err
	}
	log.Printf("%v", query)

	switch query.Table {
	case "account":
		{
			accounts, _ := dbModels.account.Get(query)
			b, err := json.Marshal(accounts)
			if err != nil {
				return nil, err
			}
			return b, nil
		}
	}
	return nil, models.UnknownError
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
			b, err := processQuery(msg, dbModels)
			if err != nil {
				log.Printf("Query error: %v", err)
				// TODO send error to client
				continue
			}
			c.Send <- b
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
