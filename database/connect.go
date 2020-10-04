package database

import (
	"fmt"
	"invoiceai/config"
	"os"

	"github.com/go-pg/pg"
	"github.com/tkanos/gonfig"
)

var DB *pg.DB

func ConnectDB() {
	dbConfig := config.Database{}

	err := gonfig.GetConf("config/config.development.json", &dbConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	pg.Connect(&pg.Options{
		User:     dbConfig.Username,
		Password: dbConfig.Password,
		Database: dbConfig.DBName,
		Addr:     dbConfig.Addr,
	})
}
