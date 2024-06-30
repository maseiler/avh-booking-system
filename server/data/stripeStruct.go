package data

import (
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"github.com/stripe/stripe-go/v78/terminal/reader"
)

func PayByCard(payment Payment) string {
	// ToDo Load Stripe Private Key from Database
	stripe.Key = "sk_test_51PBL0dCnA8pi9zjTAvmZX3sWiXme7mgL7uLZkW1yGU1Rw9DJcQvjySUShmb2y2ew76P9NmlmBFcgVHQZqEMpuzW100Rj3PgeDz"

	/*
		Register Card Reader
	*/
	// ToDo Get Location from Frontend -> LocalStorage
	// locationID := "tml_FjrpaAbsQx0TLO"
	// readerParams := &stripe.TerminalReaderParams{
	// 	Location:         stripe.String(locationID),
	// 	RegistrationCode: stripe.String("simulated-wpe"),
	// }
	// readerObj, err := reader.New(readerParams)

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

	// reader.ProcessPaymentIntent(readerObj.ID, TerminalParams)
	reader.ProcessPaymentIntent("tmr_Fjr3RwpN5Jg1xD", TerminalParams)

	// ToDo: have to capture and constantly check for Status
	// capture, err := paymentintent.Capture(pi.ID, nil)
	// if !true {
	// 	log.Fatal(err == nil)
	// }
	return pi.ID
}
