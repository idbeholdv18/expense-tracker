package bootstrap

import (
	"github/idbeholdv18/expense-tracker/internal/auth"
	"github/idbeholdv18/expense-tracker/internal/config"
	expensetypes "github/idbeholdv18/expense-tracker/internal/expense_types"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	ratelimiter "github/idbeholdv18/expense-tracker/internal/rate_limiter"
	"github/idbeholdv18/expense-tracker/internal/security/password"
	"github/idbeholdv18/expense-tracker/internal/token"
	"github/idbeholdv18/expense-tracker/internal/validation"
	"time"
)

type Services struct {
	Auth         *auth.AuthService
	Expenses     *expenses.ExpenseService
	ExpenseTypes *expensetypes.ExpenseTypesService
	Token        *token.TokenService
	RateLimiter  *ratelimiter.PostgreRateLimiter
	Validation   *validation.ValidationService
}

func RegisterServices(config *config.Config, repositories *Repositories) *Services {

	authService := &auth.AuthService{
		Repo: repositories.User,
		Hasher: &password.BcryptHasher{
			Cost: config.BcryptCost,
		},
	}

	expensesService := &expenses.ExpenseService{
		Repo: repositories.Expenses,
	}

	tokenService := &token.TokenService{
		Secret: []byte(config.JWTSecret),
	}

	expenseTypesService := &expensetypes.ExpenseTypesService{
		Repo: repositories.ExpenseTypes,
	}

	rateLimiterService := &ratelimiter.PostgreRateLimiter{
		Repo:   repositories.RateLimitLog,
		Limit:  5,
		Window: time.Second * 10,
	}

	validationService := validation.NewValidationService()

	return &Services{
		Auth:         authService,
		Expenses:     expensesService,
		ExpenseTypes: expenseTypesService,
		Token:        tokenService,
		RateLimiter:  rateLimiterService,
		Validation:   validationService,
	}
}
