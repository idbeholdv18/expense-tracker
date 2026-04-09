package httptransport

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	var appError *contract.AppError

	if errors.As(err, &appError) {
		writeJSON(w, appError.Status, appError)
		return
	}

	log.Printf(
		"error: %v | path=%s method=%s",
		err,
		r.URL.Path,
		r.Method,
	)

	writeJSON(w, http.StatusInternalServerError, &contract.AppError{
		Code:    "INTERNAL_ERROR",
		Message: "Internal server error",
		Status:  http.StatusInternalServerError,
	})
}

func ErrMethodNotAllowed() *contract.AppError {
	return &contract.AppError{
		Code:    "METHOD_NOT_ALLOWED",
		Status:  http.StatusMethodNotAllowed,
		Message: "Method not allowed",
	}
}

func ErrBadRequest() *contract.AppError {
	return &contract.AppError{
		Code:    "BAD_REQUEST",
		Status:  http.StatusBadRequest,
		Message: "Bad request",
	}
}

func ErrUnauthorized() *contract.AppError {
	return &contract.AppError{
		Code:    "UNAUTHORIZED",
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
	}
}
