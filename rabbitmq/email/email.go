package email

import (
	"fmt"
	"net/smtp"

	"github.com/go-mail/mail"
)

func Example() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("relay.alfa-bank.kz:25")
	if err != nil {
		fmt.Println("err1 = ", err.Error())
	}

	// Set the sender and recipient first
	if err := c.Mail("ocrm@alfabank.kz"); err != nil {
		fmt.Println("err2 = ", err.Error())
	}
	if err := c.Rcpt("MKapalbekov@alfabank.kz"); err != nil {
		fmt.Println("err3 = ", err.Error())
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		fmt.Println("err4 = ", err.Error())
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		fmt.Println("err5 = ", err.Error())
	}
	err = wc.Close()
	if err != nil {
		fmt.Println("err6 = ", err.Error())
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		fmt.Println("err7 = ", err.Error())
	}
}

// variables to make ExamplePlainAuth compile, without adding
// unnecessary noise there.
var (
	from       = "ocrm@alfabank.kz"
	msg        = []byte("test message")
	recipients = []string{"MKapalbekov@alfabank.kz"}
)

func ExamplePlainAuth() {
	// hostname is used by PlainAuth to validate the TLS certificate.
	hostname := "relay.alfa-bank.kz"
	//auth := smtp.PlainAuth("", "user@example.com", "password", hostname)

	err := smtp.SendMail(hostname+":25", nil /*auth*/, from, recipients, msg)
	if err != nil {
		fmt.Println("err = ", err.Error())
	}
}

func ExampleSendMail() {
	// Set up authentication information.
	//auth := smtp.PlainAuth("", "MKapalbekov@alfabank.kz", "connect!778", "VSERVER611.alfa-bank.kz")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"MKapalbekov@alfabank.kz"}
	msg := []byte("To: MKapalbekov@alfabank.kz\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	//relay.alfabank.kz
	//VSERVER611.alfa-bank.kz:25
	err := smtp.SendMail("relay.alfa-bank.kz:25", nil, "ocrm@alfabank.kz", to, msg)
	if err != nil {
		fmt.Println("err = ", err.Error())
	}
}

func LastSend() {
	fmt.Println("Try sending mail...")
	//d := mail.NewDialer("relay.alfabank.kz", 25, "ocrm@alfabank.kz", "XYZ")
	d := mail.Dialer{Host: "relay.alfabank.kz", Port: 25, StartTLSPolicy: mail.NoStartTLS}
	m := mail.NewMessage()
	m.SetHeader("From", "ocrm@alfabank.kz")
	m.SetHeader("To", "MKapalbekov@alfabank.kz")
	m.SetHeader("Subject", "Test Subj!!!")
	m.SetBody("text/plain", "This is a message body text")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed sending mail")
		panic(err)
	}
	fmt.Println("Mail sent without error")
}
