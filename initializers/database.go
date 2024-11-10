package initializers

import (
	"LibManSys/api/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn :="postgres://ohxsobfl:H543KYManb2FQMpI7mOmil6nkb5Np9Lb@rain.db.elephantsql.com/ohxsobfl"

	// container
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("------------Error connecting to database-----------------------------------")
	} else {
		log.Println("----------------------------Connected to DB----------------------------")
	}

	if err := DB.AutoMigrate(&models.Library{}); err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Library{})
	DB.AutoMigrate(&models.BookInventory{})
	DB.AutoMigrate(&models.IssueRegistry{})
	DB.AutoMigrate(&models.RequestEvents{})
	DB.AutoMigrate(&models.User{})

	log.Println("------------------------------------Migrated database-------------------------")
}
