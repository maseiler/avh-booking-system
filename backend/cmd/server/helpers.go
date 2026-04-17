package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	app.log.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// serveSPA serves static files from FrontendPath for known assets, and falls
// back to index.html for everything else so Vue Router can handle the route.
func (app *application) serveSPA(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(app.conf.FrontendPath, filepath.Clean("/"+r.URL.Path))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(app.conf.FrontendPath, "index.html"))
		return
	}
	http.ServeFile(w, r, path)
}
