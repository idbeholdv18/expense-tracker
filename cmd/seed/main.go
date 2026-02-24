package main

import (
	"context"
	"database/sql"
	"fmt"
	"github/idbeholdv18/expense-tracker/internal/auth"
	"github/idbeholdv18/expense-tracker/internal/security/password"
	"github/idbeholdv18/expense-tracker/internal/user"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db, err := sql.Open("pgx", "postgres://idbeholdv:idbeholdv@localhost:5433/expense_tracker?sslmode=disable")

	db.Exec("TRUNCATE TABLE users.users CASCADE")
	db.Exec("TRUNCATE TABLE expenses.expenses CASCADE")
	db.Exec("TRUNCATE TABLE expenses.expense_types CASCADE")

	if err != nil {
		log.Fatal("database connection err")
	}

	userRepo := user.NewUserRepository(db)

	hasher := &password.BcryptHasher{
		Cost: bcrypt.DefaultCost,
	}

	authService := &auth.AuthService{
		Repo:   userRepo,
		Hasher: hasher,
	}

	for i := range 10 {
		_, err := authService.Register(
			context.Background(),
			fmt.Sprintf("user%d@test.com", i),
			fmt.Sprintf("user%d", i),
			"password123",
		)

		if err != nil {
			log.Fatal(err)
		}
	}
}
