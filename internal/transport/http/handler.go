package httptransport

import (
	"encoding/json"
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"log"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func AppHandlerToHttpHandler(h AppHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			var appErr *contract.AppError
			if errors.As(err, &appErr) {
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
			return
		}
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
