package bootstrap

import (
	"github/idbeholdv18/expense-tracker/internal/auth"
	"github/idbeholdv18/expense-tracker/internal/domain"
	expensetypes "github/idbeholdv18/expense-tracker/internal/expense_types"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	ratelimiter "github/idbeholdv18/expense-tracker/internal/rate_limiter"
	"github/idbeholdv18/expense-tracker/internal/token"
	transport_errors "github/idbeholdv18/expense-tracker/internal/transport/http/errors"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"github/idbeholdv18/expense-tracker/internal/user"
	"github/idbeholdv18/expense-tracker/internal/validation"
)

func RegisterErrors() {
	domain.RegisterErrors(transport_errors.Register)
	auth.RegisterErrors(transport_errors.Register)
	token.RegisterErrors(transport_errors.Register)
	user.RegisterErrors(transport_errors.Register)
	expenses.RegisterErrors(transport_errors.Register)
	expensetypes.RegisterErrors(transport_errors.Register)
	middleware.RegisterJwtErrors(transport_errors.Register)
	ratelimiter.RegisterErrors(transport_errors.Register)
	validation.RegisterErrors(transport_errors.Register)
}
