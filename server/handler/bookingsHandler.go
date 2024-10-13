package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	data "github.com/maseiler/avh-booking-system/server/data"
	dbP "github.com/maseiler/avh-booking-system/server/database"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"github.com/stripe/stripe-go/v78/terminal/reader"
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
	paymentIntent := PayByCard(payment)
	response := marshalToJSON(paymentIntent, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func PayByCard(payment data.Payment) string {
	allSettings := dbP.GetSettings()
	for i := range allSettings {
		if allSettings[i].Name == "StripeAPIKey" {
			stripe.Key = allSettings[i].Value
		}
	}
	/*
		Payment Intent creation
	*/
	amountInCents := (payment.Balance * 100)
	desc := fmt.Sprintf("%s (%s) %s", payment.User.FirstName, payment.User.BierName, payment.User.LastName)
	intentParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amountInCents)),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card_present",
		}),
		CaptureMethod: stripe.String("automatic"),
		Description:   stripe.String(desc),
	}
	pi, err := paymentintent.New(intentParams)
	if !true {
		log.Fatal(err == nil)
	}

	TerminalParams := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String(pi.ID),
	}

	// tmr_Fjr3RwpN5Jg1xD
	reader.ProcessPaymentIntent(payment.CardReader, TerminalParams)
	return pi.ID
}

// Get payment intent
func ConfirmPaymentIntent(w http.ResponseWriter, r *http.Request) {
	allSettings := dbP.GetSettings()
	eMailUser := ""
	for i := range allSettings {
		if allSettings[i].Name == "StripeAPIKey" {
			stripe.Key = allSettings[i].Value
		}
		if allSettings[i].Name == "EMailUser" {
			eMailUser = allSettings[i].Value
		}
	}
	payment := UnmarshalPayment(r.Body)
	params := &stripe.TerminalReaderParams{}
	result, err := reader.Get(payment.CardReader, params)
	if err != nil {
		// Error
	}

	w.Header().Set("Content-Type", "application/json")
	if result.Action.Status != "succeeded" {
		w.Write(marshalToJSON(result, w))
		return
	}

	success := dbP.Pay(payment) // Log the payment in the Database
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(marshalToJSON(result, w))

	if !true { // ToDo: Check if sending an EMail after every payment is activated in the settings.
		//
	}
	// Send an E-Mail to the User, with all payment information included.
	receiver := payment.User.Email
	if len(receiver) == 0 {
		log.Println("Mail to " + payment.User.FirstName + " (" + payment.User.BierName + ") " + payment.User.LastName + " failed. Because no E-Mail is configured for the user.")
	} else {
		receiverCC := ""  //Normally noone needs to receive this mail in CC
		receiverBCC := "" //But it could be possible to set a BCC receiver to the one in charge with the Bank Account
		sender := eMailUser

		timeStamp := time.Now().Format("2006-01-02 15:04:05")
		// Set new User.Balance based on the payment amount.
		newBalance := payment.User.Balance - payment.Balance

		//From here on think of internationalisation or Database Settings for the texts.
		subject := "New Payment received"
		message := "<p>Hello dear " + payment.User.FirstName + " (" + payment.User.BierName + ") " + payment.User.LastName + ",<br>" +
			"We received a payment from you.</p>" +
			"<table><tbody>" +
			"<tr><td>User Balance:</td><td>" + fmt.Sprintf("%.2f", (payment.User.Balance*-1)) + "</td></tr>" +
			"<tr><td>Amount Payed:</td><td>" + fmt.Sprintf("%.2f", payment.Balance) + "</td></tr>" +
			"<tr><td>New User Balance:</td><td>" + fmt.Sprintf("%.2f", (newBalance*-1)) + "</td></tr>" +
			"<tr><td>Payment method:</td><td>" + payment.PaymentMethod + "</td></tr>" +
			"</tbody></table><br>" +
			"<p>Thank you for your payment on " + timeStamp + "</p>"
		EmailController(receiver, receiverCC, receiverBCC, sender, subject, message)
	}
}

// Cancel the current Action on the reader
func CancelReaderAction(w http.ResponseWriter, r *http.Request) {
	// Zusätzlich kann hier erst das Payment Intent gestoppt werden, bevor man die ReaderAction beendet.
	// Aktuell noch nicht nötig
	readerid := UnmarshalString(r.Body)
	params := &stripe.TerminalReaderCancelActionParams{}
	result, err := reader.CancelAction(readerid, params)
	if err != nil || result == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
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
