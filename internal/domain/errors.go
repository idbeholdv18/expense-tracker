package domain

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrMethodNotAllowed    = errors.New("method not allowed")
	ErrBadRequest          = errors.New("bad request")
	ErrInvalidReference    = errors.New("foreign key violation")
	ErrNotFound            = errors.New("not found")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
)

func RegisterErrors(register func(domainErr error, handler func(error) *contract.AppError)) {
	register(ErrMethodNotAllowed, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusMethodNotAllowed,
			"METHOD_NOT_ALLOWED",
			"method not allowed",
		)
	})

	register(ErrBadRequest, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"BAD_REQUEST",
			"bad request",
		)
	})

	register(ErrInvalidReference, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"INVALID_REFERENCE",
			"invalid referenced entity",
		)
	})

	register(ErrNotFound, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusNotFound,
			"NOT_FOUND",
			"not found",
		)
	})

	register(ErrUnauthorized, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusUnauthorized,
			"UNAUTHORIZED",
			"unauthorized",
		)
	})

	register(ErrInternalServerError, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			"internal server error",
		)
	})
}
