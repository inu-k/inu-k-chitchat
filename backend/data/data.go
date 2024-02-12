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
	var dbUser, dbPassword, DbHost, DbName, DbPort string
	if err != nil {
		fmt.Println("Error loading .env file")
		dbUser = "postgres"
		dbPassword = "postgres"
	} else {
		dbUser = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		DbHost = os.Getenv("DB_HOST")
		DbName = os.Getenv("DB_NAME")
		DbPort = os.Getenv("DB_PORT")
	}
	fmt.Println(dbUser, dbPassword, DbHost, DbName, DbPort)
	Db, err = sql.Open("postgres", fmt.Sprintf("sslmode=disable user=%s password=%s host=%s dbname=%s port=%s", dbUser, dbPassword, DbHost, DbName, DbPort))
	if err != nil {
		panic(err)
	}
	return
}
