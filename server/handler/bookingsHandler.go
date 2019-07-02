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

// GetUserDebts forwards API call to get not payed book entries
func GetUserDebts(w http.ResponseWriter, r *http.Request) {
	user := UnmarshalUser(r.Body)

	debts := dbP.GetUserDebts(user)
	response := marshalToJSON(debts, w)
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

// DeleteBookEntry forwards API call to databse to delete book entry from database
func DeleteBookEntry(w http.ResponseWriter, r *http.Request) {
	entry := UnmarshalBookEntry(r.Body)
	success := dbP.DeleteBookEntry(entry)
	validation := "ok"
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		validation = "Couldn't delete book entry."
	}
	fmt.Fprint(w, validation)
}

func userIsEmpty(user data.User) bool {
	emptyUser := data.User{ID: 0, BierName: "", FirstName: "", LastName: "", BoatName: "", Status: "", Email: "", Phone: "", Balance: 0, MaxDebt: 0}
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
