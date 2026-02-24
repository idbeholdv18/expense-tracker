package middleware

import (
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"net/http"
)

type CorsConfig struct {
	AllowedOrigin string
}

func CorsMiddleware(corsConfig *CorsConfig) func(next httptransport.AppHandler) httptransport.AppHandler {
	return func(next httptransport.AppHandler) httptransport.AppHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			w.Header().Set("Access-Control-Allow-Origin", corsConfig.AllowedOrigin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return nil
			}

			return next(w, r)
		}
	}

}
