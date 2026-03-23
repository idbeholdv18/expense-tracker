package middleware

import (
	"errors"
	ratelimiter "github/idbeholdv18/expense-tracker/internal/rate_limiter"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"net/http"
)

type KeyBuilder func(r *http.Request) (string, error)

func RateLimiterMiddleware(rl ratelimiter.RateLimiter, action string, keyBuilder KeyBuilder) func(next httptransport.AppHandler) httptransport.AppHandler {
	return func(next httptransport.AppHandler) httptransport.AppHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			keyPart, err := keyBuilder(r)

			finalKey := action + ":" + keyPart

			ok, err := rl.Allow(r.Context(), finalKey)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New(ratelimiter.ErrReachedRateLimit.Error())
			}

			return next(w, r)
		}
	}
}

func WithRateLimiter() {

}
