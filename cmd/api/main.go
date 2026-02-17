package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github/idbeholdv18/expense-tracker/internal/auth"
	"github/idbeholdv18/expense-tracker/internal/provider"
	"github/idbeholdv18/expense-tracker/internal/repository/postgres"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	port := flag.Int("port", 8080, "API server port")

	flag.Parse()

	db, err := sql.Open("pgx", "postgres://idbeholdv:idbeholdv@localhost:5433/expense_tracker?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to Postgres:", err)
	}

	fmt.Println("Connected to Postgres")

	repo := postgres.NewUserRepository(db)

	authService := auth.AuthService{
		Repo:   repo,
		Secret: []byte("secret"),
	}

	handler := provider.AuthHandler{
		Auth: &authService,
	}

	http.Handle("/login", handler.HandleLogin())
	http.Handle("/register", handler.HandleRegister())

	http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), nil)
}
