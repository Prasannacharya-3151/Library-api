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

func MigrateDB() {
	query := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	title VARCHAR(225) NOT NULL,
	author VARCHAR(225) NOT NULL,
	isbn VARCHAR(50) UNIQUE,
	total_copies INT NOT NULL DEFAULT 1,
	available_copies INT NOT NULL DEFAULT 1,
	created-at TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate books table:", err)
	}

	log.Println("books table ready")
}