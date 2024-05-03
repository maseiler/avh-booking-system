package handler

import (
	"fmt"
	"net/http"

	dataP "github.com/maseiler/avh-booking-system/server/data"
	dbP "github.com/maseiler/avh-booking-system/server/database"
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
		return "User already exists.", newUser
	}
	switch newUser.Status {
	case "Gast", "Steganleger":
		if newUser.FirstName == "" && newUser.LastName == "" {
			return "First or last name must be specified.", newUser
		} else if newUser.Email == "" {
			return "Email address must be specified.", newUser
		} else if newUser.MaxDebt == 0 {
			newUser.MaxDebt = 0 // Change if they should have Standard values during creation
		}
		return "ok", newUser
	case "Aktiv B", "Aktiv KA":
		if newUser.FirstName == "" {
			return "First name must be specified.", newUser
		} else if newUser.Email == "" {
			return "Email address must be specified.", newUser
		} else if newUser.MaxDebt == 0 {
			newUser.MaxDebt = 0 // Change if they should have Standard values during creation
		}
		return "ok", newUser
	case "AH":
		if newUser.BierName == "" {
			return "Biername must be specified.", newUser
		} else if newUser.Email == "" {
			return "Email address must be specified.", newUser
		} else if newUser.MaxDebt == 0 {
			newUser.MaxDebt = 0 // Change if they should have Standard values during creation
		}
		return "ok", newUser
	default:
		return "Select status.", newUser
	}
}

// GetUsers forwards API call to get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := [][]dataP.User{dbP.GetUsersOfColumnWithValue("status", "AH"), dbP.GetUsersOfColumnWithValue("status", "Aktiv B"), dbP.GetUsersOfColumnWithValue("status", "Aktiv KA"), dbP.GetUsersOfColumnWithValue("status", "Steganleger"), dbP.GetUsersOfColumnWithValue("status", "Gast")}
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
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
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
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	fmt.Fprint(w, validation)
}
