package database

import (
	"fmt"
	"invoiceai/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var DBConn  is a global variable we use to make connections
var (
	DBConn *gorm.DB
)

// InitDatabase is a function we use to init the database
func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	fmt.Println("Database Connected")
	DBConn.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
