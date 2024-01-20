package handler

import (
	"fmt"
	"net/http"

	dataP "github.com/maseiler/avh-booking-system/server/data"
	dbP "github.com/maseiler/avh-booking-system/server/database"
)

// GetAllItems forwards API call to get all items from database
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := dbP.GetAllItems()
	response := marshalToJSON(items, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// AddItem forwards API call to add new item to database
func AddItem(w http.ResponseWriter, r *http.Request) {
	newItem := UnmarshalItem(r.Body)

	if dbP.NewItemExists(newItem) {
		w.WriteHeader(http.StatusBadRequest)
		// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
		fmt.Fprint(w, "Item already exists")
		return
	}

	validation, newItem := validateNewItemArguments(newItem)
	if validation == "ok" {
		dbP.AddItem(newItem)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}

// ModifyItem forwards API call to replaces all values of an item
func ModifyItem(w http.ResponseWriter, r *http.Request) {
	item := UnmarshalItem(r.Body)

	validation, newItem := validateNewItemArguments(item)

	if validation == "ok" {
		dbP.ModifyItem(newItem)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}

// validates input parameters when creating an new item
func validateNewItemArguments(newItem dataP.Item) (validation string, item dataP.Item) {
	if newItem.Name == "" {
		return "Name must be specified.", newItem
	} else if newItem.Price == 0 {
		return "Price must be specified.", newItem
	} else if newItem.Size == 0 {
		return "Size must be specified.", newItem
	} else if newItem.Unit == "" {
		return "Unit must be specified.", newItem
	} else if newItem.Type == "" {
		return "Type must be specified.", newItem
	} else {
		return "ok", newItem
	}
	//no check for Enabled because its default value is true
}

// DeleteItem forwards API call to delete item from database
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	item := UnmarshalItem(r.Body)

	validation := ""
	if !dbP.ItemExists(item) {
		validation = "Item doesn't exist."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		dbP.DeleteItem(item)
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	}
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	fmt.Fprint(w, validation)
}
