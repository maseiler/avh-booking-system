package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	// set up file server
	fileServer := http.FileServer(http.Dir(app.conf.FrontendPath))

	// set up routes
	mux := http.NewServeMux()

	// HTTP handler
	mux.Handle("GET /{$}", fileServer)
	//mux.HandleFunc("GET /account/{id}", app.getAccount)

	// WebSocket handler
	mux.Handle("/ws", app.wsHandler)

	return app.logRequest(app.securityHeaders(mux))
}
