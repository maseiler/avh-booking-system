package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	dataP "../data"
	dbP "../database"
)

// AddUser forwards API call to add new user to database
func AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newUser dataP.User
	err := decoder.Decode(&newUser)

	if err != nil {
		panic(err)
	}

	dbP.AddUser(newUser)
}

// GetUsers forwards API call to get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := dbP.GetUsers()
	response := marshalToJSON(users, w)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// func AddItem(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)

// 	var newItem dataP.Item
// 	err := decoder.Decode(&newItem)

// 	if err != nil {
// 		panic(err)
// 	}

// 	dbP.AddItem(newItem)
// }

// GetItems forwards API call to get all items from database
func GetItems(w http.ResponseWriter, r *http.Request) {
	items := dbP.GetItems()
	response := marshalToJSON(items, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Checkout TODO
func Checkout(w http.ResponseWriter, r *http.Request) {
	//sql query
	fmt.Fprintf(w, "Check this out")
}
