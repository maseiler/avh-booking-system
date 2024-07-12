package handler

import (
	"net/http"

	dbP "github.com/maseiler/avh-booking-system/server/database"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/terminal/reader"
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

func GetStripeCardReader(w http.ResponseWriter, r *http.Request) {
	allSettings := dbP.GetSettings()
	for i := range allSettings {
		if allSettings[i].Name == "StripeAPIKey" {
			stripe.Key = allSettings[i].Value
		}
	}

	readParams := &stripe.TerminalReaderListParams{}
	// readParams.Limit = stripe.Int64(10)
	result := reader.List(readParams)

	allReadersJSON := marshalToJSON(result.List(), w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(allReadersJSON)
}
