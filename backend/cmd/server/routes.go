package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	// set up routes
	mux := http.NewServeMux()

	// WebSocket handler
	mux.Handle("/ws", app.wsHandler)

	// SPA catch-all: serve index.html for any path not matched above,
	// letting Vue Router handle client-side navigation.
	mux.HandleFunc("/", app.serveSPA)

	return app.logRequest(app.securityHeaders(mux))
}
