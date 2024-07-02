package handler

import (
	"fmt"
	"net/http"

	data "github.com/maseiler/avh-booking-system/server/data"
	dbP "github.com/maseiler/avh-booking-system/server/database"
	"github.com/stripe/stripe-go/v78"

	"github.com/stripe/stripe-go/v78/terminal/reader"
	// "github.com/stripe/stripe-go#beta-sdks"
)

// GetLastNBookEntries forwards API call to get the n latest bookings from database
func GetLastNBookEntries(w http.ResponseWriter, r *http.Request) {
	n := UnmarshalInt(r.Body)
	bookings := dbP.GetLastNBookEntries(n)
	response := marshalToJSON(bookings, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetBookEntriesFromUserBetween forwards API call to get book entries of user within time span
func GetBookEntriesFromUserBetween(w http.ResponseWriter, r *http.Request) {
	userFromTo := UnmarshalUserFromTo(r.Body)
	from := ConvertStringToTime(userFromTo.From)
	to := ConvertStringToTime(userFromTo.To)

	bookings := dbP.GetBookEntriesOfUserBetween(userFromTo.User, from, to)

	response := marshalToJSON(bookings, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetBookEntriesFromItemBetween forwards API call to get book entries of item within time span
func GetBookEntriesFromItemBetween(w http.ResponseWriter, r *http.Request) {
	itemFromTo := UnmarshalItemFromTo(r.Body)
	from := ConvertStringToTime(itemFromTo.From)
	to := ConvertStringToTime(itemFromTo.To)

	bookings := dbP.GetBookEntriesOfItemBetween(itemFromTo.Item, from, to)

	response := marshalToJSON(bookings, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetPaymentsOfUser forwards API call to get payments of user
func GetPaymentsOfUser(w http.ResponseWriter, r *http.Request) {
	userFromTo := UnmarshalUserFromTo(r.Body)
	from := ConvertStringToTime(userFromTo.From)
	to := ConvertStringToTime(userFromTo.To)

	bookings := dbP.GetPaymentsOfUser(userFromTo.User, from, to)

	response := marshalToJSON(bookings, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetUserDebts forwards API call to get unpaid book entries
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
		// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
		fmt.Fprint(w, "Select an Item.")
		return
	}
	if userIsEmpty(cart.User) {
		w.WriteHeader(http.StatusBadRequest)
		// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
		fmt.Fprint(w, "Select a User.")
		return
	}
	success := dbP.Checkout(cart)
	validation := "ok"
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
		validation = "Reached max balance."
	}
	fmt.Fprint(w, validation)
}

// Pay forwards API call to databse to pay current balance
func Pay(w http.ResponseWriter, r *http.Request) {
	payment := UnmarshalPayment(r.Body)
	paymentIntent := data.PayByCard(payment)
	response := marshalToJSON(paymentIntent, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	return

	validation := ""
	// for cardPayment.Status != "succeeded" {
	// 	// wait until succeeded
	// }
	// if cardPayment.Status == "succeeded" {
	// 	validation = "Could not pay by card."
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, validation)
	// 	return
	// }
	success := dbP.Pay(payment)
	if success {
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	} else {
		validation = "Couldn't pay."
		w.WriteHeader(http.StatusBadRequest)
	}
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	fmt.Fprint(w, validation)
}

// Get payment intent
func ConfirmPaymentIntent(w http.ResponseWriter, r *http.Request) {
	// ToDo Load Stripe Private Key from Database
	stripe.Key = "sk_test_51PBL0dCnA8pi9zjTAvmZX3sWiXme7mgL7uLZkW1yGU1Rw9DJcQvjySUShmb2y2ew76P9NmlmBFcgVHQZqEMpuzW100Rj3PgeDz"

	payment := UnmarshalPayment(r.Body)
	// params := &stripe.PaymentIntentConfirmParams{}
	// result, err := paymentintent.Confirm(payment.IntentID, params)

	// params := &stripe.TerminalReaderConfirmPaymentIntentParams{
	// 	PaymentIntent: stripe.String(payment.IntentID),
	// }
	// result, err := reader.ConfirmPaymentIntent("tmr_Fjr3RwpN5Jg1xD", params)

	params := &stripe.TerminalReaderParams{}
	result, err := reader.Get("tmr_Fjr3RwpN5Jg1xD", params)

	if err != nil {
		// Error
	}

	w.Header().Set("Content-Type", "application/json")
	if result.Action.Status != "succeeded" {
		w.Write(marshalToJSON(result, w))
		return
	}
	// validation := ""
	success := dbP.Pay(payment)
	if success {
		// validation = "ok"
		w.WriteHeader(http.StatusOK)
	} else {
		// validation = "Couldn't pay."
		w.WriteHeader(http.StatusBadRequest)
	}
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	// ToDo: switch from fmt.Fprint to w.Write and give json response for frontend to figure out the right message in correct language
	w.Write(marshalToJSON(result, w))
	// fmt.Fprint(w, validation)
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
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	fmt.Fprint(w, validation)
}

// UndoBookEntry creates a new book entry with inversed balance and adjusts the user's balance accordingly.
func UndoBookEntry(w http.ResponseWriter, r *http.Request) {
	entry := UnmarshalBookEntry(r.Body)
	success := dbP.UndoBookEntry(entry)
	validation := "ok"
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		validation = "Couldn't undo book entry."
	}
	// ToDo: internaitonalize this message - maybe send only error-codes and do the text at client side
	fmt.Fprint(w, validation)
}

func userIsEmpty(user data.User) bool {
	emptyUser := data.User{ID: 0, BierName: "", FirstName: "", LastName: "", BoatName: "", Status: "", Email: "", Phone: "", Balance: 0, MaxDebt: 0}
	return user == emptyUser
}

func itemsAreEmpty(items []data.CartItem) bool {
	return len(items) == 0
}
