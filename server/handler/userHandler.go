package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	data "github.com/maseiler/avh-booking-system/server/data"
	db "github.com/maseiler/avh-booking-system/server/database"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := db.GetUsersByQuery("SELECT * FROM users;")
	usersAsJson, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(usersAsJson)
}

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser data.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &newUser)

	err = validateUserArguments(&newUser)

	if err == nil {
		fmt.Println("")
		db.AddNewUser(newUser)
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
	}
}

// checks if necessary fields are set and sets max debt
func validateUserArguments(newUser *data.User) error {
	if db.NewUserExists(*newUser) {
		return errors.New(fmt.Sprintf("user already exists: %+v", newUser))
	}
	switch newUser.Status {
	case "Gast", "Steganleger":
		if newUser.FirstName == "" && newUser.LastName == "" {
			return errors.New("first or last name not specified")
		} else if newUser.Email == "" && newUser.Phone == "" {
			return errors.New("email address or phone number not specified")
		} else if newUser.MaxDebt == 0 {
			newUser.MaxDebt = -50
		}
		return nil
	case "Aktiv B", "Aktiv KA":
		if newUser.FirstName == "" {
			return errors.New("first name not specified")
		} else if newUser.MaxDebt == 0 {
			newUser.MaxDebt = -50
		}
		return nil
	case "AH":
		if newUser.BierName == "" {
			return errors.New("biername not specified")
		} else if newUser.MaxDebt == 0 {
			newUser.MaxDebt = -100
		}
		return nil
	default:
		return errors.New("invalid status")
	}
}
