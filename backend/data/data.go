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
	if err != nil {
		panic(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	Db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=chitchat sslmode=disable", dbUser, dbPassword))
	if err != nil {
		panic(err)
	}
	return
}
