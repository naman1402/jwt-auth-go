package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var e error
	dsn := os.Getenv("DB")
	// Using ORM library
	DB, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic("database not connected")
	}

}
