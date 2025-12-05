package main

import (
	"encoding/json"
	"fmt"
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
	Table  string   `json:"table"`
	Filter []Filter `json:"filter"`
}

type Filter struct {
	Column   string `json:"column"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

func getQueryFromMessage(data interface{}) (*Query, error) {
	// Convert the interface{} back to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	// Unmarshal directly into the Query struct
	var query Query
	if err := json.Unmarshal(jsonBytes, &query); err != nil {
		return nil, fmt.Errorf("failed to unmarshal query: %w", err)
	}
	/*
		m := data.(map[string]interface{})
		query := Query{}
		if table, ok := m["table"].(string); ok {
			query.Table = table
		} else {
			return nil, models.BadJson
		}
	*/
	/*
		if _, ok := m["filter"]; ok {
			filters, ok := m["filter"].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("not a map[string]interface{}")
			}
			filter := &Filter{
				Column:   filters["column"].(string),
				Operator: filters["operator"].(string),
				Value:    filters["value"].(string),
			}
			query.Filter = *filter
		} else {
			return nil, models.BadJson
		}
	*/

	/*
		slice, ok := m["filter"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("not a slice")
		}

		filters := make([]Filter, 0, len(slice))

		for idx, item := range slice {
			// Convert each item to a map
			f, ok := item.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("item at index %d is not a map", idx)
			}

			filter := Filter{
				Operator: f["operator"].(string),
				Column:   f["column"].(string),
				Value:    f["value"].(string),
			}
			filters = append(filters, filter)
		}

		fmt.Println(filters)
		query.Filter = filters
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
func processQuery(message Message, dbModels *dbModels) ([]byte, error) {
	query, err := getQueryFromMessage(message.Payload)
	if err != nil {
		return nil, err
	}
	log.Printf("%v", query)

	switch query.Table {
	case "accounts":
		{
			accounts, _ := dbModels.account.GetAll()
			for _, account := range accounts {
				log.Println(account)
			}
			/*
				account, err := dbModels.account.Get(1)
				if err != nil {
					return nil, err
				}

				b, err := json.Marshal(account)
				if err != nil {
					return nil, err
				}
				return b, nil
			*/
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
			/*
				query := getQueryFromMessage(msg.Payload)
				log.Printf("%v", query.Table)
				account, err := dbModels.account.Get(1)
				if err != nil {
					if errors.Is(err, models.ErrNoRecord) {
						panic("not found")
					} else {
						panic(err)
					}

					return
				}
				fmt.Printf("%v", account)

				b, err := json.Marshal(account)
				if err != nil {
					fmt.Println(err)
					return
				}
			*/
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
