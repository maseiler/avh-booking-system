package main

import (
	"encoding/json"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"log"
)

type WebSocketService struct {
	hub       *models.Hub
	validator *validation.WebSocketValidator
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

func getMutationFromMessage(data interface{}) (*models.Mutation, *models.WsError) {
	// Convert the interface{} back to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal interface"}
	}

	// Unmarshal directly into the Mutation struct
	var mutation models.Mutation
	if err := json.Unmarshal(jsonBytes, &mutation); err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal to Mutation"}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(mutation)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadStruct, Message: err.Error(), Details: "Invalid struct"}
	}

	return &mutation, nil
}

func NewWebSocketService() *WebSocketService {
	hub := &models.Hub{
		Clients:    make(map[*models.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
	}

	go hub.Run()

	return &WebSocketService{hub: hub, validator: validation.NewWebSocketValidator()}
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
	case models.TableAccount:
		{
			accounts, _ := dbModels.account.Get(query)
			msg := models.Message{
				Type:    models.MsgTypeResponse,
				Payload: map[string]interface{}{"accounts": accounts},
			}
			b, err := json.Marshal(msg)
			if err != nil {
				return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal response including Account"}
			}
			return b, nil
		}
	case models.TableProduct:
		// ...
	}

	return nil, &models.WsError{Code: models.WsUnknown, Message: "Unknown error", Details: "Could not process query"}
}

// processMutation unmarshals the message and initiates a database mutation
func processMutation(message models.Message, dbModels *dbModels) ([]byte, *models.WsError) {

	mutation, err := getMutationFromMessage(message.Payload)
	if err != nil {
		return nil, err
	}

	switch mutation.Operation {
	case models.OpInsert:
		{
			switch mutation.Table {
			case models.TableAccount:
				jsonBytes, err := json.Marshal(mutation.Values)
				if err != nil {
					return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal interface"}
				}

				// Convert the interface{} back to JSON bytes
				account := models.Account{}
				err = json.Unmarshal(jsonBytes, &account)
				if err != nil {
					return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal JSON"}
				}
				row, err := dbModels.account.Insert(account)
				if err != nil {
					return nil, &models.WsError{Code: models.WsInternalError, Message: err.Error(), Details: "Could not create account"}
				}
				msg := models.Message{
					Type:    models.MsgTypeResponse,
					Payload: map[string]interface{}{"row": row},
				}
				b, err := json.Marshal(msg)
				if err != nil {
					return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal response"}
				}
				return b, nil
			}
		}
	case models.OpUpdate:
		{
			return nil, &models.WsError{Code: models.WsUnknown, Message: "Not yet implemented", Details: "TBD"}
		}
	case models.OpDelete:
		{
			return nil, &models.WsError{Code: models.WsUnknown, Message: "Not yet implemented", Details: "TBD"}
		}
	}

	return nil, &models.WsError{Code: models.WsUnknown, Message: "Unknown error", Details: "Could not process mutation"}
}

func sendError(c *models.Client, wsErr *models.WsError) {
	log.Printf("%v", wsErr)
	msg := models.Message{
		Type: models.MsgTypeError,
		Payload: map[string]interface{}{
			"error": wsErr},
	}
	errBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	c.Send <- errBytes
}

// ReadPump pumps messages from the websocket connection to the hub
func ReadPump(c *models.Client, dbModels *dbModels) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512 * 1024) // 512KB max message size

	wsValidator := validation.NewWebSocketValidator()

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

		err = wsValidator.ValidateMessage(message)
		if err != nil {
			wsErr := &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
			sendError(c, wsErr)
		}

		// unmarshal message
		msg := models.Message{}
		if err := json.Unmarshal(message, &msg); err != nil {
			wsErr := &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal message"}
			sendError(c, wsErr)
			continue
		}

		// Handle different message types
		switch msg.Type {
		case models.MsgTypeBroadcast:
			c.Hub.Broadcast <- message
		case models.MsgTypePing:
			response := models.Message{Type: models.MsgTypePong}
			pong, _ := json.Marshal(response)
			c.Send <- pong
		case models.MsgTypeQuery:

			// Convert the interface{} back to JSON bytes (Query)
			jsonBytes, err := json.Marshal(msg.Payload)
			if err != nil {
				panic(err)
			}

			err = wsValidator.ValidateQuery(jsonBytes)
			if err != nil {
				panic(err)
			}

			// --------------------------------------

			b, wsErr := processQuery(msg, dbModels)
			if wsErr != nil {
				sendError(c, wsErr)
				continue
			}
			c.Send <- b
		case models.MsgTypeMutation:
			response, wsErr := processMutation(msg, dbModels)
			if wsErr != nil {
				sendError(c, wsErr)
				continue
			}
			c.Send <- response
		default:
			log.Fatalf("Unknown message type: %v", msg.Type)
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
