package bootstrap

import (
	"github/idbeholdv18/expense-tracker/internal/auth"
	expensetypes "github/idbeholdv18/expense-tracker/internal/expense_types"
	"github/idbeholdv18/expense-tracker/internal/expenses"
)

type Handlers struct {
	Auth         *auth.AuthHandler
	Expenses     *expenses.ExpenseHandler
	ExpenseTypes *expensetypes.ExpenseTypesHandler
}

func RegisterHandlers(services *Services) *Handlers {
	authHandler := &auth.AuthHandler{
		Auth:       services.Auth,
		Token:      services.Token,
		Validation: services.Validation,
		Email:      services.Email,
	}

	expensesHandler := &expenses.ExpenseHandler{
		Service: services.Expenses,
	}

	expenseTypesHandler := &expensetypes.ExpenseTypesHandler{
		Service: services.ExpenseTypes,
	}

	return &Handlers{
		Auth:         authHandler,
		Expenses:     expensesHandler,
		ExpenseTypes: expenseTypesHandler,
	}
}
