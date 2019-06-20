package handler

import (
	"fmt"
	"net/http"

	dbP "../database"
)

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
