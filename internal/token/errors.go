package token

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
)

func RegisterErrors(register func(domainErr error, handler func(error) *contract.AppError)) {

	register(ErrInvalidSigningMethod, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"BAD_REQUEST",
			"invalid token signing method",
		)
	})

	register(ErrInvalidToken, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusUnauthorized,
			"INVALID_TOKEN",
			"invalid or expired token",
		)
	})
}
