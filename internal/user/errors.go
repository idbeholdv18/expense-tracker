package user

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

func RegisterErrors(register func(domainErr error, handler func(error) *contract.AppError)) {
	register(ErrUserNotFound, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusNotFound,
			"USER_NOT_FOUND",
			"User not found",
		)
	})

	register(ErrUserAlreadyExists, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusConflict,
			"USER_ALREADY_EXISTS",
			"User already exists",
		)
	})
}
