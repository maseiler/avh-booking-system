package handler

import (
	"fmt"
	"net/http"

	data "../data"
	dbP "../database"
)

// GetLastNBookings forwards API call to get the n latest bookings from database
func GetLastNBookings(w http.ResponseWriter, r *http.Request) {
	n := UnmarshalInt(r.Body)

	bookings := dbP.GetLastNBookings(n)
	response := marshalToJSON(bookings, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Checkout forwards API call to add Cart content into database
func Checkout(w http.ResponseWriter, r *http.Request) {
	cart := UnmarshalCart(r.Body)
	if itemsAreEmpty(cart.CartItems) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Select an Item.")
		return
	}
	if userIsEmpty(cart.User) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Select a User.")
		return
	}
	success := dbP.Checkout(cart)
	validation := "ok"
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		validation = "Reached max balance."
	}
	fmt.Fprint(w, validation)
}

// Pay forwards API call to databse to pay current balance
func Pay(w http.ResponseWriter, r *http.Request) {
	user := UnmarshalUser(r.Body)
	success := dbP.Pay(user)
	validation := ""
	if success {
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	} else {
		validation = "Couldn't pay."
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}

func userIsEmpty(user data.User) bool {
	emptyUser := data.User{UserID: 0, BierName: "", FirstName: "", LastName: "", Status: "", Email: "", Balance: 0, Phone: "", MaxDebt: 0}
	if user == emptyUser {
		return true
	}
	return false
}

func itemsAreEmpty(items []data.CartItem) bool {
	if len(items) == 0 {
		return true
	}
	return false
}
