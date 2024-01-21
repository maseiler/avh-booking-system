package handler

import (
	"net/http"

	dbP "github.com/maseiler/avh-booking-system/server/database"
)

// GetSettings forwards API call to get all settings from database
func GetSettings(w http.ResponseWriter, r *http.Request) {
	settingMap := dbP.GetSettings()
	response := marshalToJSON(settingMap, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateSetting forwards API call to update setting values to database
func UpdateSetting(w http.ResponseWriter, r *http.Request) {
	setting := UnmarshalSetting(r.Body)
	dbP.UpdateSetting(setting)
	w.WriteHeader(http.StatusOK)
}
