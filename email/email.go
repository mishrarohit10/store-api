package email

import (
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

func SendEmail(to string, subject string, body string, ch chan string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "sender@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.example.com", 587, "username", "password")

	var err error

	for i := 0; i < 3; i++ {
		if err = d.DialAndSend(m); err == nil {
			ch <- "Email sent to " + to
			return
		}
		time.Sleep(5 * time.Second) // Retry after 5 seconds
	}
	ch <- "Failed to send email to " + to + ": " + err.Error()
}

func Mail() {

	recipients := []struct {
		Email string
		Subject string
		Body string
	}{
		{"recipient1@example.com", "Hello from Golang", "This is the first email."},
		{"recipient2@example.com", "Greetings from Go", "This is the second email."},
		// Add more recipients here
	}

	emailStatus := make(chan string)

	for _, r := range recipients {
		go SendEmail(r.Email, r.Subject, r.Body, emailStatus)
	}

	for range recipients {
		status := <-emailStatus
		log.Println(status)
	}
}
