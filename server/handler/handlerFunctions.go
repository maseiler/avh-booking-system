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

	validation := validateInputArguments(newUser)

	if validation == "ok" {
		dbP.AddUser(newUser)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}

// validates input parameters when creating a new user
func validateInputArguments(newUser dataP.User) (validation string) {
	if dbP.UserExists(newUser) {
		return "User already exists"
	}
	switch newUser.Status {
	case "Gast":
		if newUser.FirstName == "" && newUser.LastName == "" {
			return "First or last name must be specified."
		} else if newUser.Email == "" && newUser.Phone == "" {
			return "Email address or phone number must be specified"
		}
		return "ok"
	case "Aktiv B", "Aktiv KA", "AH":
		if newUser.BierName == "" && newUser.FirstName == "" {
			return "Biername or first name must be specified"
		}
		return "ok"
	default:
		return "Select status"
	}
}

// GetUsers forwards API call to get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := [][]dataP.User{dbP.GetUsersOfColumnWithValue("Status", "AH"), dbP.GetUsersOfColumnWithValue("Status", "Aktiv B"), dbP.GetUsersOfColumnWithValue("Status", "Aktiv KA"), dbP.GetUsersOfColumnWithValue("Status", "Gast")}
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
