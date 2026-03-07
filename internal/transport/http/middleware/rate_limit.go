package middleware

import (
	ratelimiter "github/idbeholdv18/expense-tracker/internal/rate_limiter"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"net/http"
)

type TokenGenerator func(*http.Request) string

func RateLimiterMiddleware(rl ratelimiter.RateLimiter, action string, tokenGenerator TokenGenerator) func(next httptransport.AppHandler) httptransport.AppHandler {
	return func(next httptransport.AppHandler) httptransport.AppHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			token := tokenGenerator(r)

			ok, err := rl.Allow(r.Context(), action, token)
			if err != nil || !ok {
				return err
			}

			return next(w, r)
		}
	}
}
