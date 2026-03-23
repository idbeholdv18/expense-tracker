package main

import (
	"embed"
	"fmt"

	"github/idbeholdv18/expense-tracker/internal/config"
	"github/idbeholdv18/expense-tracker/internal/database"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	cfg := config.Load()
	db := database.New(cfg.DatabaseURL)

	goose.SetBaseFS(embedMigrations)

	entries, _ := embedMigrations.ReadDir("migrations")
	for _, e := range entries {
		fmt.Println("Migration file:", e.Name())
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	log.Println("Applying migrations...")
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	log.Println("Migrations applied successfully!")
}
