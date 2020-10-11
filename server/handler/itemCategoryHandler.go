package handler

import (
	"fmt"
	"net/http"

	dbP "github.com/maseiler/avh-booking-system/server/database"
)

// GetItemCategories forwards API call to get all item categories from database
func GetItemCategories(w http.ResponseWriter, r *http.Request) {
	itemCategories := dbP.GetItemCategories()
	response := marshalToJSON(itemCategories, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// AddItemCategory forwards API call to add new item category into database
func AddItemCategory(w http.ResponseWriter, r *http.Request) {
	newItemCategory := UnmarshalItemCategory(r.Body)

	if dbP.ItemCategoryExists(newItemCategory) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Item category already exists")
		return
	}

	dbP.AddItemCategory(newItemCategory)
	w.WriteHeader(http.StatusOK)
}

// TODO rename

// DeleteItemCategory forwards API call to delete item category from database
func DeleteItemCategory(w http.ResponseWriter, r *http.Request) {
	itemCategory := UnmarshalItemCategory(r.Body)

	validation := ""
	if !dbP.ItemCategoryExists(itemCategory) {
		validation = "Item categorey doesn't exist."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		dbP.DeleteItemCategory(itemCategory)
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, validation)
}
