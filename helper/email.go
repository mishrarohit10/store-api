package helper

import (
	"LibManSys/api/initializers"
	"LibManSys/api/models"
	"log"
	"gopkg.in/gomail.v2"

)

type recipient struct {
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

	dialer := gomail.Dialer{Host: "smtp.gmail.com", Port: 587, Username: "mishrarohit316x@gmail.com", Password: "xxxxxxxxxxxxxx"}
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

	SendToMultipleRecipientsWithGomailV2(recipients, "A new book is added to the inventory", "<h1>A New Book is added in the library check it now</h1>")
}
