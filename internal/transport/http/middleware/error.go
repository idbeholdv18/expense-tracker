package middleware

import (
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	transport_errors "github/idbeholdv18/expense-tracker/internal/transport/http/errors"
	"net/http"
)

func ErrorMiddleware() func(next httptransport.AppHandler) httptransport.AppHandler {
	return func(next httptransport.AppHandler) httptransport.AppHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			err := next(w, r)
			if err == nil {
				return nil
			}

			if handler, ok := transport_errors.Get(err); ok {
				appErr := handler(err)
				return appErr
			}
			return err
		}
	}
}
