package handler

import (
	"fmt"
	"net/http"

	dbP "github.com/maseiler/avh-booking-system/server/database"
)

// Login confirms/denies entered password
func Login(w http.ResponseWriter, r *http.Request) {
	pw := UnmarshalString(r.Body)
	pws := dbP.GetPasswords()
	success := validatePasswords(pws, pw)
	var validation string
	if success {
		w.WriteHeader(http.StatusOK)
		validation = "Confirmed"
	} else {
		w.WriteHeader(http.StatusBadRequest)
		validation = "Denied"
	}
	w.Header().Set("Content-Type", "application/json")
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	fmt.Fprint(w, validation)
}

func validatePasswords(pws []string, pw string) bool {
	for i := 0; i < len(pws); i++ {
		if pws[i] == pw {
			return true
		}
	}
	return false
}

// ChangeAdminPassword checks if old password is valid and forwards API call to database
func ChangeAdminPassword(w http.ResponseWriter, r *http.Request) {
	oldAndNewPassword := UnmarshalStringArray(r.Body)
	currentPasswords := dbP.GetPasswords()

	oldPasswordValid := false
	for i := 0; i < len(currentPasswords); i++ {
		if oldAndNewPassword[0] == currentPasswords[i] {
			oldPasswordValid = true
		}
	}

	validation := ""
	if oldPasswordValid {
		dbP.ModifyPassword(oldAndNewPassword[0], oldAndNewPassword[1])
		w.WriteHeader(http.StatusOK)
		validation = "Changed Password."
	} else {
		w.WriteHeader(http.StatusBadRequest)
		validation = "Old password is not valid."
	}
	w.Header().Set("Content-Type", "application/json")
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	fmt.Fprint(w, validation)
}
