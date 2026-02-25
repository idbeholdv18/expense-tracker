package main

import (
	"context"
	"fmt"
	"github/idbeholdv18/expense-tracker/internal/bootstrap"
	"github/idbeholdv18/expense-tracker/internal/config"
	"github/idbeholdv18/expense-tracker/internal/database"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	"log"
	"math/rand"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.Load()
	db := database.New(cfg.DatabaseURL)

	if _, err := db.Exec("TRUNCATE TABLE users.users CASCADE"); err != nil {
		log.Fatal("Seed err:", err)
	}
	if _, err := db.Exec("TRUNCATE TABLE expenses.expenses CASCADE"); err != nil {
		log.Fatal("Seed err:", err)
	}
	if _, err := db.Exec("TRUNCATE TABLE expenses.expense_types CASCADE"); err != nil {
		log.Fatal("Seed err:", err)
	}

	repositories := bootstrap.RegisterRepositories(db)

	services := bootstrap.RegisterServices(cfg, repositories)

	ctx := context.Background()

	rand.Seed(time.Now().UnixNano())

	for i := range 10 {
		user, err := services.Auth.Register(
			ctx,
			fmt.Sprintf("user%d@test.com", i),
			fmt.Sprintf("user%d", i),
			"password123",
		)

		if err != nil {
			log.Fatal("Seed err during user creating:", err)
		}

		for k := range 4 {
			et, err := services.ExpenseTypes.Create(
				ctx,
				user.ID,
				fmt.Sprintf("expense type %d", k),
			)

			if err != nil {
				log.Fatal("Seed err during exepnse type creating:", err)
			}

			for j := range rand.Intn(5) {
				candidate := &expenses.CreateExpenseInput{
					Amount:        float64(rand.Intn(100)),
					ExpenseTypeID: et.ID,
					Currency:      "RUB",
					Description:   fmt.Sprintf("expense %s", j),
					ExpenseDate:   time.Now(),
				}
				_, err := services.Expenses.Create(
					ctx,
					user.ID,
					candidate,
				)

				if err != nil {
					log.Fatal("Seed err during exepnse creating:", err)
				}

			}
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}
