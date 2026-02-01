package ws

import (
	"encoding/json"
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/gorilla/websocket"
	"log/slog"
)

// ReadPump pumps messages from the ws connection to the hub
func (s *Service) ReadPump(c *models.Client) {
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
				s.log.Error("WebSocket error: %v", err.Error())
			}
			break
		}

		// Validate JSON schema
		err = s.validator.ValidateMessage(message)
		if err != nil {
			wsErr := &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
			s.sendError(c, wsErr)
		}

		// Unmarshal message
		msg := models.Message{}
		if err := json.Unmarshal(message, &msg); err != nil {
			wsErr := &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal message"}
			s.sendError(c, wsErr)
			continue
		}

		s.log.Debug("Received message", slog.String("msg", fmt.Sprintf("%v", msg)))

		// Route
		switch msg.Type {
		case models.MsgTypeBroadcast:
			c.Hub.Broadcast <- message

		case models.MsgTypePing:
			b, wsErr := s.processPing(msg)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}
			c.Send <- b

		case models.MsgTypeQuery:
			b, wsErr := s.processQuery(msg)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}
			c.Send <- b

		case models.MsgTypeMutation:
			mutation, wsErr := unmarshalInterface[models.Mutation](msg.Payload)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}

			response, id, wsErr := s.processMutation(mutation)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}
			c.Send <- response

			message, wsErr = s.prepareBroadcast(mutation, id)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}

			c.Hub.Broadcast <- message

		default:
			s.log.Error("Unknown message type:", slog.String("msg.Type", msg.Type.String()))
		}
	}
}

// WritePump pumps messages from the hub to the ws connection
func (s *Service) WritePump(c *models.Client) {
	defer c.Conn.Close()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
