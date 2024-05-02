package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/maseiler/avh-booking-system/server/data"
	dbP "github.com/maseiler/avh-booking-system/server/database"
)

func SendTestMail(w http.ResponseWriter, r *http.Request) {
	data.EmailController("dev@leon-ebert.de", "", "", os.Getenv("AVHBS_EMAIL_USER"), "Test E-Mail", "Dies ist eine Test Nachricht vom AVHBS Server")
}

func SendCurrentDebts(w http.ResponseWriter, r *http.Request) {
	until := UnmarshalString(r.Body)
	if until == "" {
		until = time.Now().String()[:10]
	}
	receiver := "dev@leon-ebert.de"         //Change to Database Setting E-Mail
	receiverCC := ""                        //Change to Database Setting E-Mail
	receiverBCC := ""                       //Change to Database Setting E-Mail
	sender := os.Getenv("AVHBS_EMAIL_USER") //Change to Database Setting E-Mail

	//From here on think of internationalisation or Database Settings for the texts.
	subject := "Current Debts until " + until
	message := "<p>Hello dear Friend,<br>" +
		"Unitl this day: " + until + " the current debts per category look like this:</p>" +
		"<table><thead><tr><th>Category</th><th>Debt</th></tr></thead><tbody>"
	for category, debt := range dbP.GetAllCategoryDebts(until) {
		message += "<tr><td>" + category + "</td><td>" + fmt.Sprintf("%.2f", debt) + "</td></tr>"
	}
	message += "</tbody></table>" +
		"<p>In this list, positive values mean debt while negative values represent credit.</p>"

	data.EmailController(receiver, receiverCC, receiverBCC, sender, subject, message)
}

// Brauche noch Hilfe mit der Datenbankfunktion
// Später könnte dieser Handler z.B. über cronjobs automatisch jede Montag-Nacht aufgerufen werden.
func SendLastWeeksBookings(w http.ResponseWriter, r *http.Request) {
	// receiver := "dev@leon-ebert.de"         //Change to Database Setting E-Mail
	// receiverCC := ""                        //Change to Database Setting E-Mail
	// receiverBCC := ""                       //Change to Database Setting E-Mail
	// sender := os.Getenv("AVHBS_EMAIL_USER") //Change to Database Setting E-Mail

	// //From here on think of internationalisation or Database Settings for the texts.
	// subject := "Current Debts until " + until
	// message := "<p>Hello dear Friend,<br>" +
	// 	"In the last week the following items had been purchased:</p>" +
	// 	"<table><thead><tr><th>Item</th><th>Size</th><th>Unit</th><th>Amount</th></tr></thead><tbody>"
	// for item, size, unit, amount := range dbP.GetLasWeekSoldItems(){
	// 	message += "<tr><td>" + item + "</td><td>" + size + "</td><td>" + unit + "</td><td>" + amount + "</td></tr>"
	// }
	// message += "</tbody></table>"
	// data.EmailController(receiver, receiverCC, receiverBCC, sender, subject, message)
}
