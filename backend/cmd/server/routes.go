package main

import (
	"net/http"
	"os"
)

func (app *application) routes() http.Handler {
	// set up file server
	staticPath := os.Getenv("AVHBS_FRONTEND_PATH")
	fileServer := http.FileServer(http.Dir(staticPath))

	// set up routes
	mux := http.NewServeMux()

	// HTTP handler
	mux.Handle("GET /{$}", fileServer)
	mux.HandleFunc("GET /account/{id}", app.getAccount)

	// WebSocket handler
	wsHandler := NewWebSocketHandler(app.wsService)
	mux.HandleFunc("/ws", wsHandler.HandleConnections)

	return app.logRequest(app.securityHeaders(mux))
}
