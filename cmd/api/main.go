package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github/idbeholdv18/expense-tracker/internal/auth"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	"github/idbeholdv18/expense-tracker/internal/middleware"
	"github/idbeholdv18/expense-tracker/internal/provider"
	"github/idbeholdv18/expense-tracker/internal/repository/postgres"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
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

	userRepo := postgres.NewUserRepository(db)
	expensesRepo := postgres.NewExpensesRepository(db)

	authService := auth.AuthService{
		Repo:   userRepo,
		Secret: []byte("secret"),
	}

	expensesService := expenses.ExpenseService{
		Repo: expensesRepo,
	}

	authHandler := provider.AuthHandler{
		Auth: &authService,
	}

	expensesHandler := provider.ExpenseHandler{
		Service: &expensesService,
	}

	jwtMiddleware := middleware.JwtMiddleware([]byte("secret"))

	http.Handle("/api/v1/login", authHandler.HandleLogin())
	http.Handle("/api/v1/register", authHandler.HandleRegister())
	http.Handle("/api/v1/expenses", jwtMiddleware(expensesHandler.HandleExpense()))

	http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), nil)
}
