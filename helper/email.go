package helper

import (
	"LibManSys/api/initializers"
	"LibManSys/api/models"
	"log"

	"gopkg.in/gomail.v2"
	// "time"
)

// func SendEmail(to string, subject string, body string, ch chan string) {

// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "boomer.com")
// 	m.SetHeader("To", to)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/plain", body)

// 	d := gomail.NewDialer("smtp.example.com", 587, "username", "password")

// 	var err error

// 	for i := 0; i < 3; i++ {
// 		if err = d.DialAndSend(m); err == nil {
// 			ch <- "Email sent to " + to
// 			return
// 		}
// 		time.Sleep(5 * time.Second) // Retry after 5 seconds
// 	}
// 	ch <- "Failed to send email to " + to + ": " + err.Error()
// }

// func Mail() {

// 	recipients := []struct {
// 		Email string
// 		Subject string
// 		Body string
// 	}{
// 		{"recipient1@example.com", "Hello from Golang", "This is the first email."},
// 		{"recipient2@example.com", "Greetings from Go", "This is the second email."},
// 		// Add more recipients here
// 	}

// 	emailStatus := make(chan string)

// 	for _, r := range recipients {
// 		go SendEmail(r.Email, r.Subject, r.Body, emailStatus)
// 	}

// 	for range recipients {
// 		status := <-emailStatus
// 		log.Println(status)
// 	}
// }

type recipient struct {
	// Name    string `json:"Name"`
	Address string `json:"email"`
}

func SendToMultipleRecipientsWithGomailV2(recipients []recipient, subject, message string) {

	log.Println("inside mailing")

	mail := gomail.NewMessage()

	addresses := make([]string, len(recipients))
	for i, recipient := range recipients {
		addresses[i] = mail.FormatAddress(recipient.Address, "user")
	}

	log.Println(addresses, "adressssss---------------------------")


	mail.SetHeader("From", "mishrarohit316x@gmail.com")
	mail.SetHeader("To", addresses...)
	// mail.SetHeader("To", "rohit.mishra@xenonstack.com")
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", message)

	dialer := gomail.Dialer{Host: "smtp.gmail.com", Port: 587, Username: "mishrarohit316x@gmail.com", Password: "rnmz qvsx cysp ylol"}
	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
	mail.Reset()
}

func Mail() {

	var users []models.User

	results := initializers.DB.Find(&users)

	log.Println("start--------------------------------------------")

	log.Println(users, "------------------------------------------------------------")

	log.Println("end--------------------------------------------")
	if results.Error != nil {
		log.Println("error finding users")
		return
	}

	log.Println(len(users), "length of []")

	var recipients []recipient

	for i := 0; i < len(users); i++ {
		recipients = append(recipients, recipient{users[i].Email})
	}

	log.Println(recipients, "this is recipientts")

	// recipients := []recipient{{
	// 	Name:    "User 1",
	// 	Address: "user1@gmail.com",
	// }, {
	// 	Name:    "User 2",
	// 	Address: "user2@yandex.com",
	// }}

	SendToMultipleRecipientsWithGomailV2(recipients, "A new book is added to the inventory", "<h1>A New Book is added in the library check it now</h1>")
}
