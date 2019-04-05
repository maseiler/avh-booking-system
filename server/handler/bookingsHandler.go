package handler

import (
	"fmt"
	"net/http"

	dbP "../database"
)

// GetLastNBookings forwards API call to get the n latest bookings from database
func GetLastNBookings(w http.ResponseWriter, r *http.Request) {
	n := UnmarshalInt(r.Body)

	bookings := dbP.GetLastNBookings(n)
	response := marshalToJSON(bookings, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Checkout forwards API call to add Cart content into database
func Checkout(w http.ResponseWriter, r *http.Request) {
	cart := UnmarshalCart(r.Body)

	success := dbP.Checkout(cart)
	validation := "ok"
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		validation = "Reached max balance."
	}
	fmt.Fprint(w, validation)
}
