package handler

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	data "github.com/maseiler/avh-booking-system/server/data"
	dbP "github.com/maseiler/avh-booking-system/server/database"
)

func SendTestMail(w http.ResponseWriter, r *http.Request) {
	EmailController("dev@leon-ebert.de", "", "", os.Getenv("AVHBS_EMAIL_USER"), "Test E-Mail", "Dies ist eine Test Nachricht vom AVHBS Server")
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

	EmailController(receiver, receiverCC, receiverBCC, sender, subject, message)
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

// BuildMessage ...
func BuildMessage(mail *data.Mail) string {
	settingMap := dbP.GetSettings()
	eMailSender := ""
	for _, set := range settingMap {
		if set.Name == "EMailSender" {
			eMailSender = set.Value
		}
	}
	header := ""
	header += fmt.Sprintf("From: %s <%s>\r\n", eMailSender, mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += fmt.Sprintf("Content-Type: text/html; charset=UTF-8\r\n")
	header += "\r\n" + mail.Body

	return header
}

func send(mailObject data.Mail) {
	mail := data.Mail{}
	mail.Sender = mailObject.Sender
	mail.To = mailObject.To
	mail.Subject = mailObject.Subject
	mail.Body = mailObject.Body

	messageBody := BuildMessage(&mail)

	settingMap := dbP.GetSettings()
	eMailHost := ""
	eMailPort := ""
	eMailPass := ""
	for _, set := range settingMap {
		if set.Name == "EMailHost" {
			eMailHost = set.Value
		}
		if set.Name == "EMailPort" {
			eMailPort = set.Value
		}
		if set.Name == "EMailPass" {
			eMailPass = set.Value
		}
	}

	// smtpServer := SMTPServer{Host: os.Getenv("AVHBS_EMAIL_HOST"), Port: os.Getenv("AVHBS_EMAIL_PORT")}
	smtpServer := data.SMTPServer{Host: eMailHost, Port: eMailPort}
	smtpServer.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpServer.Host,
	}

	auth := smtp.PlainAuth("", mail.Sender, eMailPass, smtpServer.Host)

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.TLSConfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		log.Panic(err)
	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	wr, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = wr.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = wr.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()
	log.Println("Mail sent successfully")
}

// EmailController controls email operations
func EmailController(to, cc, bcc, sender, subject string, message string) {

	toArr := strings.Split(to, ";")
	ccArr := strings.Split(cc, ";")
	bccArr := strings.Split(bcc, ";")
	body := message
	body += "<br><br><p>Please contact <i>Brenner</i> or <i>Käptn</i> if something went wrong.<br>" +
		"Sincerely yours,<br>" +
		"<i>(A)kademischer (V)erein (H)ütte (B)uchungs(s)ystem Server ;)</i></p>"
	request := data.Mail{
		Sender:  sender,
		To:      toArr,
		Cc:      ccArr,
		Bcc:     bccArr,
		Subject: subject,
		Body:    body,
	}
	// fmt.Println("request: ", request)

	send(request)
}
