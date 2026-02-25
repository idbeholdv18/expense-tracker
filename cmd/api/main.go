package main

import (
	"github/idbeholdv18/expense-tracker/internal/bootstrap"
	"github/idbeholdv18/expense-tracker/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.Load()
	bootstrap.New(cfg).Run()
}
