package bootstrap

import (
	"database/sql"
	expensetypes "github/idbeholdv18/expense-tracker/internal/expense_types"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	ratelimiter "github/idbeholdv18/expense-tracker/internal/rate_limiter"
	"github/idbeholdv18/expense-tracker/internal/user"
)

type Repositories struct {
	User         user.UserRepository
	Expenses     expenses.ExpenseRepository
	ExpenseTypes expensetypes.ExpenseTypeRepository
	RateLimitLog ratelimiter.RateLimitLogRepository
}

func RegisterRepositories(db *sql.DB) *Repositories {
	userRepo := user.NewUserRepository(db)
	expensesRepo := expenses.NewExpensesRepository(db)
	expenseTypesRepo := expensetypes.NewExepenseTypesRepository(db)
	rateLimitLogRepo := ratelimiter.NewLimitLogRepository(db)

	return &Repositories{
		User:         userRepo,
		Expenses:     expensesRepo,
		ExpenseTypes: expenseTypesRepo,
		RateLimitLog: rateLimitLogRepo,
	}
}
