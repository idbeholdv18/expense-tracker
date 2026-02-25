package bootstrap

import (
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"net/http"
)

func registerRoutes(middlewares *Middlewares, handlers *Handlers) {

	authHandler := middleware.Apply(
		handlers.Auth.HandleLogin(),
		middlewares.CORS,
		middlewares.Errors,
	)
	http.Handle("/api/v1/login", httptransport.AppHandlerToHttpHandler(authHandler))

	registerHandler := middleware.Apply(
		handlers.Auth.HandleRegister(),
		middlewares.CORS,
		middlewares.Errors,
	)
	http.Handle("/api/v1/register", httptransport.AppHandlerToHttpHandler(registerHandler))

	expensesHandler := middleware.Apply(
		handlers.Expenses.HandleExpense(),
		middlewares.JWT,
		middlewares.CORS,
		middlewares.Errors,
	)
	http.Handle("/api/v1/expenses", httptransport.AppHandlerToHttpHandler(expensesHandler))

	expenseTypesHandler := middleware.Apply(
		handlers.ExpenseTypes.HandleExpenseTypes(),
		middlewares.JWT,
		middlewares.CORS,
		middlewares.Errors,
	)
	http.Handle("/api/v1/expense-types", httptransport.AppHandlerToHttpHandler(expenseTypesHandler))
}
