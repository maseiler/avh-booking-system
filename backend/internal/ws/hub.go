package ws

import (
	"encoding/json"
	"log/slog"
)

// Hub maintains active clients and broadcasts messages
type Hub struct {
	log        *slog.Logger
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.log.Info("Client registered", slog.String("client.ID", client.ID))

			// Notify others about new client
			msg := Message{Type: MsgTypeRegister, Payload: map[string]string{"id": client.ID}}
			notification, err := json.Marshal(msg)
			if err != nil {
				h.log.Error("failed to marshal register notification", slog.String("error", err.Error()))
				break
			}
			h.broadcastMessage(notification, client)

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)

				h.log.Info("Client unregistered", slog.String("client.ID", client.ID))

				// Notify others about disconnection
				msg := Message{Type: MsgTypeUnRegister, Payload: map[string]string{"id": client.ID}}
				notification, err := json.Marshal(msg)
				if err != nil {
					h.log.Error("failed to marshal unregister notification", slog.String("error", err.Error()))
					break
				}
				h.broadcastMessage(notification, nil)
			}

		case message := <-h.Broadcast:
			h.broadcastMessage(message, nil)
		}
	}
}

func (h *Hub) broadcastMessage(message []byte, exclude *Client) {
	var drop []*Client

	for client := range h.Clients {
		if client != exclude {
			select {
			case client.Send <- message:
			default:
				drop = append(drop, client)
			}
		}
	}

	for _, client := range drop {
		if _, ok := h.Clients[client]; ok {
			delete(h.Clients, client)
			close(client.Send)
		}
	}
}
