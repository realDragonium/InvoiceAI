package database

import (
	"fmt"
	"invoiceai/model"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/tkanos/gonfig"
)

var DB *pg.DB

func ConnectDB() {
	dbConfig := model.DbConfig{}

	err := gonfig.GetConf("config/config.real.development.json", &dbConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	DB = pg.Connect(&pg.Options{
		User:     dbConfig.Username,
		Password: dbConfig.Password,
		Database: dbConfig.DBName,
		Addr:     dbConfig.Addr,
	})
	setupBasicDB(DB)
}

func setupBasicDB(db *pg.DB) {
	if _, err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"pgcrypto\";"); err != nil {
		panic(err)
	}

	err := createTables(db)
	if err != nil {
		panic(err)
	}

}

// createSchema creates database schema for User and Story models.
func createTables(db *pg.DB) error {
	models := []interface{}{
		(*model.DBUser)(nil),
		(*model.NewUser)(nil),
		(*model.User)(nil),
	}

	for _, model := range models {
		err := (*db).Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        true,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
