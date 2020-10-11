package handler

import (
	"fmt"
	"net/http"

	dbP "github.com/maseiler/avh-booking-system/server/database"
)

// GetItemCategoryMaps forwards API call to get all ItemCategoryMaps from database
func GetItemCategoryMaps(w http.ResponseWriter, r *http.Request) {
	itemCategoryMaps := dbP.GetItemCategoryMaps()
	response := marshalToJSON(itemCategoryMaps, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// AddItemCategoryMap forwards API call to add new item category map into database
func AddItemCategoryMap(w http.ResponseWriter, r *http.Request) {
	icMap := UnmarshalItemCategoryMap(r.Body)

	if dbP.ItemCategoryMapExists(icMap) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Item category map already exists")
		return
	}

	dbP.AddItemCategoryMap(icMap)
	w.WriteHeader(http.StatusOK)
}

// TODO rename

// DeleteItemCategoryMap forwards API call to delete item category map from database
func DeleteItemCategoryMap(w http.ResponseWriter, r *http.Request) {
	icMap := UnmarshalItemCategoryMap(r.Body)

	validation := ""
	if !dbP.ItemCategoryMapExists(icMap) {
		validation = "Item categorey map doesn't exist."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		dbP.DeleteItemCategoryMap(icMap)
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, validation)
}
