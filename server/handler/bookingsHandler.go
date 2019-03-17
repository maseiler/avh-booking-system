package handler

import (
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
