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

	validation, newUser := validateUserArguments(newUser)

	if validation == "ok" {
		dbP.AddUser(newUser)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}

// validates input parameters when creating an new user
func validateUserArguments(newUser dataP.User) (validation string, user dataP.User) {
	if dbP.NewUserExists(newUser) {
		return "User already exists", newUser
	}
	switch newUser.Status {
	case "Gast":
		if newUser.FirstName == "" && newUser.LastName == "" {
			return "First or last name must be specified.", newUser
		} else if newUser.Email == "" && newUser.Phone == "" {
			return "Email address or phone number must be specified", newUser
		} else if newUser.MaxDebt == 0 {
			newUser.MaxDebt = 50 //TODO
			return "ok", newUser
		}
		return "ok", newUser
	case "Aktiv B", "Aktiv KA", "AH":
		if newUser.BierName == "" && newUser.FirstName == "" {
			return "Biername or first name must be specified", newUser
		} else if newUser.Status == "Aktiv B" || newUser.Status == "Aktiv KA" {
			newUser.MaxDebt = 50
		} else if newUser.Status == "AH" {
			newUser.MaxDebt = 100
		}
		return "ok", newUser
	default:
		return "Select status", newUser
	}
}

// GetUsers forwards API call to get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := [][]dataP.User{dbP.GetUsersOfColumnWithValue("Status", "AH"), dbP.GetUsersOfColumnWithValue("Status", "Aktiv B"), dbP.GetUsersOfColumnWithValue("Status", "Aktiv KA"), dbP.GetUsersOfColumnWithValue("Status", "Gast")}
	response := marshalToJSON(users, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// ModifyUser forwards API call to replaces all values of an user
func ModifyUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user dataP.User
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	validation := ""
	if !dbP.UserExists(user) {
		validation = "User doesn't exist."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		dbP.ModifyUser(user)
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, validation)
}

// DeleteUser forwards API call to delete user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user dataP.User
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	validation := ""
	if !dbP.UserExists(user) {
		validation = "User doesn't exist."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		dbP.DeleteUser(user)
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, validation)
}

// GetItems forwards API call to get all items from database
func GetItems(w http.ResponseWriter, r *http.Request) {
	items := dbP.GetItems()
	response := marshalToJSON(items, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// AddItem forwards API call to add new item to database
func AddItem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newItem dataP.Item
	err := decoder.Decode(&newItem)

	if err != nil {
		panic(err)
	}

	if dbP.NewItemExists(newItem) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Item already exists")
		return
	}

	validation, newItem := validateItemArguments(newItem)
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
	decoder := json.NewDecoder(r.Body)

	var item dataP.Item
	err := decoder.Decode(&item)

	if err != nil {
		panic(err)
	}

	validation, newItem := validateItemArguments(item)

	if validation == "ok" {
		dbP.ModifyItem(newItem)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}

// validates input parameters when creating an new item
func validateItemArguments(newItem dataP.Item) (validation string, item dataP.Item) {
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
}

// DeleteItem forwards API call to delete item from database
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var item dataP.Item
	err := decoder.Decode(&item)

	if err != nil {
		panic(err)
	}

	validation := ""
	if !dbP.ItemExists(item) {
		validation = "Item doesn't exist."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		dbP.DeleteItem(item)
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, validation)
}

// GetLastBookings forwards API call to get the 50 latest bookings from database
func GetLastBookings(w http.ResponseWriter, r *http.Request) {
	bookings := dbP.GetLastBookings()
	response := marshalToJSON(bookings, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Checkout TODO
func Checkout(w http.ResponseWriter, r *http.Request) {
	//sql query
	fmt.Fprintf(w, "Check this out")
}
