package bootstrap

import (
	"github/idbeholdv18/expense-tracker/internal/auth"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, services *Services, middlewares *Middlewares, handlers *Handlers) {
	loginHandler := middleware.Apply(
		handlers.Auth.HandleLogin(),
		middlewares.CORS,
		middleware.RateLimiterMiddleware(services.RateLimiter, "login", auth.LoginRateLimiterKeyBuilder),
		middlewares.Errors,
	)
	registerHandler := middleware.Apply(
		handlers.Auth.HandleRegister(),
		middlewares.CORS,
		middleware.RateLimiterMiddleware(services.RateLimiter, "register", auth.RegisterRateLimiterKeyBuilder),
		middlewares.Errors,
	)
	mux.Handle("/api/v1/auth/login", httptransport.AppHandlerToHttpHandler(loginHandler))
	mux.Handle("/api/v1/auth/register", httptransport.AppHandlerToHttpHandler(registerHandler))

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
	mux.Handle("/api/v1/expense-types/", httptransport.AppHandlerToHttpHandler(expenseTypesHandler))
	mux.Handle("/api/v1/expense-types", httptransport.AppHandlerToHttpHandler(expenseTypesHandler))
}
