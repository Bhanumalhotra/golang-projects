package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect for GORM
)

var (
	db *gorm.DB // Create a global variable to hold the database connection
)

// Connect to the database
func Connect() {
	// Create the database connection with the desired configuration
	d, err := gorm.Open("root@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d

}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
