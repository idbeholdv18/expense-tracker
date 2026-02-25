package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(databaseURL string) *sql.DB {
	db, err := sql.Open("pgx", databaseURL)

	if err != nil {
		log.Fatal("Failed to open database:", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to Postgres:", err)
	}

	fmt.Println("Connected to Postgres")

	return db
}
