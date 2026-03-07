package ratelimiter

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrReachedRateLimit = errors.New("reached rate limit")
)

func RegisterErrors(register func(error, func(error) *contract.AppError)) {
	register(ErrReachedRateLimit, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"REACHED RATE LIMIT",
			"reached rate limit",
		)
	})
}
