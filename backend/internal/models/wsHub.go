package models

import (
	"encoding/json"
	"log"
	"sync"
)

// Hub maintains active clients and broadcasts messages
type Hub struct {
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
			log.Printf("Client %s registered", client.ID)

			// Notify others about new client
			notification, _ := json.Marshal(map[string]string{
				"type": "user_joined",
				"id":   client.ID,
			})
			h.broadcastMessage(notification, client)

		case client := <-h.Unregister:
			h.Mu.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				h.Mu.Unlock()

				log.Printf("Client %s unregistered", client.ID)

				// Notify others about disconnection
				notification, _ := json.Marshal(map[string]string{
					"type": "user_left",
					"id":   client.ID,
				})
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
	h.Mu.RLock()
	defer h.Mu.RUnlock()

	for client := range h.Clients {
		if client != exclude {
			select {
			case client.Send <- message:
			default:
				// Client's send channel is full, close it
				delete(h.Clients, client)
				close(client.Send)
			}
		}
	}
}
