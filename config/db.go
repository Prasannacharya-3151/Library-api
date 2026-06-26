package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq" // driver only — we never call it directly, hence the underscore import
)

var DB *sql.DB

func ConnectDB() {
	host:= os.Getenv("DB_HOST")
	port:= os.Getenv("DB_PORT")
	user:= os.Getenv("DB_USER")
	password:= os.Getenv("DB_PASSWORD")
	dbname:= os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disabel",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Faile to ping a db", err)
	}

	log.Printf("Databse connected successfully")
	DB = db
}