package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(databaseURL string) *sql.DB {
	var db *sql.DB
	var err error

	for i := 0; i < 15; i++ {
		db, err = sql.Open("pgx", databaseURL)
		if err == nil && db.Ping() == nil {
			break
		}
		log.Println("Waiting for DB to be ready...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	log.Println("DB connected! Starting API...")

	return db
}
