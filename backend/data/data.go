package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	var dbUser, dbPassword string
	if err != nil {
		fmt.Println("Error loading .env file")
		dbUser = "postgres"
		dbPassword = "postgres"
	} else {
		dbUser = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
	}
	Db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=chitchat sslmode=disable", dbUser, dbPassword))
	if err != nil {
		panic(err)
	}
	return
}
