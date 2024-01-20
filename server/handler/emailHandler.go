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

	dbP "github.com/maseiler/avh-booking-system/server/database"
)

// Mail ....
type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

// SMTPServer ...
type SMTPServer struct {
	Host      string
	Port      string
	TLSConfig *tls.Config
}

// ServerName ...
func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

// EmailController controls email operations
func EmailController(to, cc, bcc, sender, subject string, message string) {

	toArr := strings.Split(to, ";")
	ccArr := strings.Split(cc, ";")
	bccArr := strings.Split(bcc, ";")
	body := message
	request := Mail{
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

// BuildMessage ...
func (mail *Mail) BuildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s <%s>\r\n", os.Getenv("AVHBS_EMAIL_NAME"), mail.Sender)
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

func send(mailObject Mail) {
	mail := Mail{}
	mail.Sender = mailObject.Sender
	mail.To = mailObject.To
	mail.Subject = mailObject.Subject
	mail.Body = mailObject.Body

	messageBody := mail.BuildMessage()

	smtpServer := SMTPServer{Host: os.Getenv("AVHBS_EMAIL_HOST"), Port: os.Getenv("AVHBS_EMAIL_PORT")}
	smtpServer.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpServer.Host,
	}

	auth := smtp.PlainAuth("", mail.Sender, os.Getenv("AVHBS_EMAIL_PASS"), smtpServer.Host)

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
		"<p>In this list, positive values mean debt while negative values represent credit.</p>" +
		"<p>Please contact <i>Brenner</i> or <i>Käptn</i> if something went wrong.<br>" +
		"Sincerely yours,<br>" +
		"<i>(A)kademischer (V)erein (H)ütte (B)uchungs(s)ystem Server ;)</i></p>"

	EmailController(receiver, receiverCC, receiverBCC, sender, subject, message)
}
