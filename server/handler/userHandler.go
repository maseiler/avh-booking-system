package handler

import (
	"fmt"
	"net/http"

	dataP "../data"
	dbP "../database"
)

// AddUser forwards API call to add new user to database
func AddUser(w http.ResponseWriter, r *http.Request) {
	newUser := UnmarshalUser(r.Body)

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
	user := UnmarshalUser(r.Body)

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
	user := UnmarshalUser(r.Body)

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
