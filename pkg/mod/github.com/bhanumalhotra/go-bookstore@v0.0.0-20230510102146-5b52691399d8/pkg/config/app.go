package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect for GORM
)

var (
	db *gorm.DB // Create a global variable to hold the database connection
)

// Connect to the database
func Connect() {
	var err error

	// Create the database connection with the desired configuration
	db, err = gorm.Open("mysql", "bhanu:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Print a message indicating that the database connection was successful
	log.Println("Connected to database")

	// Migrate the schema to the database
	db.AutoMigrate(&Book{})
	log.Println("Database Migrated")
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
