package auth

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
)

func RegisterErrors(register func(domainErr error, handler func(error) *contract.AppError)) {

	register(ErrUserAlreadyExists, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusConflict,
			"USER_ALREADY_EXISTS",
			"User already exists",
		)
	})

	register(ErrInvalidCredentials, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusUnauthorized,
			"INVALID_CREDENTIALS",
			"Invalid credentials",
		)
	})

	register(ErrUnauthorized, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusUnauthorized,
			"UNAUTHORIZED",
			"Unauthorized",
		)
	})

}
