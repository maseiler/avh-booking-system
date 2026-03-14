package ws

import (
	"encoding/json"
	"log/slog"
	"sync"
)

// Hub maintains active clients and broadcasts messages
type Hub struct {
	Log        *slog.Logger
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Mu         sync.RWMutex
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mu.Lock()
			h.Clients[client] = true
			h.Mu.Unlock()
			h.Log.Info("Client registered", slog.String("client.ID", client.ID))

			// Notify others about new client
			msg := Message{Type: MsgTypeRegister, Payload: map[string]string{"id": client.ID}}
			notification, err := json.Marshal(msg)
			if err != nil {
				h.Log.Error("failed to marshal register notification", slog.String("error", err.Error()))
				break
			}
			h.broadcastMessage(notification, client)

		case client := <-h.Unregister:
			h.Mu.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				h.Mu.Unlock()

				h.Log.Info("Client unregistered", slog.String("client.ID", client.ID))

				// Notify others about disconnection
				msg := Message{Type: MsgTypeUnRegister, Payload: map[string]string{"id": client.ID}}
				notification, err := json.Marshal(msg)
				if err != nil {
					h.Log.Error("failed to marshal unregister notification", slog.String("error", err.Error()))
					break
				}
				h.broadcastMessage(notification, nil)
			} else {
				h.Mu.Unlock()
			}

		case message := <-h.Broadcast:
			h.broadcastMessage(message, nil)
		}
	}
}

func (h *Hub) broadcastMessage(message []byte, exclude *Client) {
	var drop []*Client

	h.Mu.RLock()
	for client := range h.Clients {
		if client != exclude {
			select {
			case client.Send <- message:
			default:
				drop = append(drop, client)
			}
		}
	}
	h.Mu.RUnlock()

	if len(drop) > 0 {
		h.Mu.Lock()
		for _, client := range drop {
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		}
		h.Mu.Unlock()
	}
}
