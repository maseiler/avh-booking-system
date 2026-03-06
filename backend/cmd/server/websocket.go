package main

import (
	"github.com/av-huette/avh-booking-system/internal/ws"
)

/*
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(service *ws.Service, dbModels *DbModels) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}

		client := service.NewClient(conn)
		service.Hub().Register <- client

		go service.WritePump(client)
		go service.ReadPump(client, dbModels)
	}
}
*/

type WebSocketHandler struct {
	app     *application
	Service *ws.Service
}

func NewWebSocketHandler(app *application, stores ws.Stores) *WebSocketHandler {
	return &WebSocketHandler{
		app:     app,
		Service: ws.NewService(stores, app.log),
	}
}
