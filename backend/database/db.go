package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DBConn  is a global variable we use to make connections
var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dsn := "user=invoice_ai password=super_secret dbname=invoice_ai port=5432"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	fmt.Println("Database Connected")

}
