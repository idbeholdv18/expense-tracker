package middleware

import (
	"encoding/json"
	"github/idbeholdv18/expense-tracker/internal/contract"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	transport_errors "github/idbeholdv18/expense-tracker/internal/transport/http/errors"
	"log"
	"net/http"
)

func ErrorMiddleware(next httptransport.AppHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		if handler, ok := transport_errors.Get(err); ok {
			appErr := handler(err)
			writeJSON(w, appErr.Status, appErr)
			return
		}

		log.Printf(
			"internal error: %v | path=%s method=%s",
			err,
			r.URL.Path,
			r.Method,
		)

		writeJSON(w, http.StatusInternalServerError, &contract.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_ERROR",
			Message: "Internal server error",
		})
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
