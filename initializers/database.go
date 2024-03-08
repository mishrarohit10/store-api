package initializers

import (
	"LibManSys/api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "os"
)


var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn :="postgres://ohxsobfl:H543KYManb2FQMpI7mOmil6nkb5Np9Lb@rain.db.elephantsql.com/ohxsobfl"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("------------Error connecting to database-----------------------------------")
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