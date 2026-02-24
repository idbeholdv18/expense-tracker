package middleware

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

func RegisterJwtErrors(register func(error, func(error) *contract.AppError)) {
	register(ErrUnauthorized, func(err error) *contract.AppError {
		return &contract.AppError{
			Status:  http.StatusUnauthorized,
			Code:    "UNAUTHORIZED",
			Message: "Unauthorized",
		}
	})
}
