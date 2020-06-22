package models

import (
	"github.com/jinzhu/gorm"
	// In order to connect to a database, you need to import its driver first.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB global var
var DB *gorm.DB

// ConnectToDB : connection to database func
func ConnectToDB() {
	/*
		Create a new connection by defining which
		kind of db we use and how to access it
	*/
	database, err := gorm.Open("sqlite3", "bookStore.db")

	if err != nil {
		panic("Faild to connect to DB")
	}

	database.AutoMigrate(&Book{})

	// populate the the DB variable with our database instance
	DB = database
}
