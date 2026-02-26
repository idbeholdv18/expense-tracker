package bootstrap

import (
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, middlewares *Middlewares, handlers *Handlers) {
	authHandler := middleware.Apply(
		handlers.Auth.HandleAuth(),
		middlewares.CORS,
		middlewares.Errors,
	)
	mux.Handle("/api/v1/auth/", httptransport.AppHandlerToHttpHandler(authHandler))

	expensesHandler := middleware.Apply(
		handlers.Expenses.HandleExpense(),
		middlewares.JWT,
		middlewares.CORS,
		middlewares.Errors,
	)
	mux.Handle("/api/v1/expenses/", httptransport.AppHandlerToHttpHandler(expensesHandler))
	mux.Handle("/api/v1/expenses", httptransport.AppHandlerToHttpHandler(expensesHandler))

	expenseTypesHandler := middleware.Apply(
		handlers.ExpenseTypes.HandleExpenseTypes(),
		middlewares.JWT,
		middlewares.CORS,
		middlewares.Errors,
	)
	mux.Handle("/api/v1/expense-types", httptransport.AppHandlerToHttpHandler(expenseTypesHandler))
}
