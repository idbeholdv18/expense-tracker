package bootstrap

import (
	"database/sql"
	expensetypes "github/idbeholdv18/expense-tracker/internal/expense_types"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	"github/idbeholdv18/expense-tracker/internal/user"
)

type Repositories struct {
	User         user.UserRepository
	Expenses     expenses.ExpenseRepository
	ExpenseTypes expensetypes.ExpenseTypeRepository
}

func RegisterRepositories(db *sql.DB) *Repositories {
	userRepo := user.NewUserRepository(db)
	expensesRepo := expenses.NewExpensesRepository(db)
	expenseTypesRepo := expensetypes.NewExepenseTypesRepository(db)

	return &Repositories{
		User:         userRepo,
		Expenses:     expensesRepo,
		ExpenseTypes: expenseTypesRepo,
	}
}
