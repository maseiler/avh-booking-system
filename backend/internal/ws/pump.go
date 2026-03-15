package ws

import (
	"encoding/json"
	"log/slog"

	"github.com/gorilla/websocket"
)

// ReadPump pumps messages from the ws connection to the hub
func (s *Service) ReadPump(c *Client) {
	defer func() {
		c.Cancel()
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
				s.log.Error("WebSocket error", slog.String("error", err.Error()))
			}
			break
		}

		// Validate JSON schema
		err = s.validator.validate(message)
		if err != nil {
			wsErr := &WSError{Code: WSBadJSON, Message: err.Error(), Details: "JSON does not comply with Message schema"}
			s.sendError(c, wsErr)
			continue
		}

		// Unmarshal message
		msg := Message{}
		if err := json.Unmarshal(message, &msg); err != nil {
			wsErr := &WSError{Code: WSBadJSON, Message: err.Error(), Details: "Could not unmarshal message"}
			s.sendError(c, wsErr)
			continue
		}

		s.log.Debug("Received message", slog.Any("msg", msg))

		// Route
		switch msg.Type {
		case MsgTypeBroadcast:
			c.Hub.Broadcast <- message

		case MsgTypePing:
			b, wsErr := s.processPing(msg)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}
			c.Send <- b

		case MsgTypeQuery:
			b, wsErr := s.processQuery(c.Ctx, msg)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}
			c.Send <- b

		case MsgTypeMutation:
			mutation, wsErr := unmarshalInterface[Mutation](msg.Payload)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}

			response, id, wsErr := s.processMutation(c.Ctx, mutation)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}
			c.Send <- response

			message, wsErr = s.prepareBroadcast(c.Ctx, mutation, id)
			if wsErr != nil {
				s.sendError(c, wsErr)
				continue
			}

			c.Hub.Broadcast <- message

		default:
			s.log.Error("Unknown message type", slog.String("msg_type", msg.Type.String()))
		}
	}
}

// WritePump pumps messages from the hub to the ws connection
func (s *Service) WritePump(c *Client) {
	defer c.Conn.Close()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
