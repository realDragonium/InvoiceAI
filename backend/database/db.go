package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

// var DBConn  is a global variable we use to make connections
var (
	DB *pgx.Conn
)

func InitDatabase() {
	var err error
	DB, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

}
