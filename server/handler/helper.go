package handler

import (
	"encoding/json"
	"net/http"
)

func marshalToJSON(object interface{}, w http.ResponseWriter) (response []byte) {
	response, err := json.Marshal(object)
	handleHTTPError(err, w)
	return response
}

func handleHTTPError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
