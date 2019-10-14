package handler

import (
	"net/http"

	dbP "github.com/maseiler/avh-booking-system/server/database"
)

// GetBookingStats forwards API call to perfromq query to return a map of the time stamp and total number of bought drinks.
func GetBookingStats(w http.ResponseWriter, r *http.Request) {
	days := UnmarshalInt(r.Body)
	stats := dbP.GetBookingStats(days)
	json := marshalToJSON(stats, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// GetFavoriteItemsStats forwards API call to perform query to return a map of item IDs and it's total number bought (if greater than 0)
func GetFavoriteItemsStats(w http.ResponseWriter, r *http.Request) {
	stats := dbP.GetFavoriteItemsStats()
	json := marshalToJSON(stats, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
