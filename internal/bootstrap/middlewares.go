package bootstrap

import (
	"github/idbeholdv18/expense-tracker/internal/config"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"net/http"
)

type Middlewares struct {
	CORS      middleware.Middleware
	JWT       middleware.Middleware
	Errors    middleware.Middleware
	RateLimit middleware.Middleware
}

func RegisterMiddlewares(config *config.Config, services *Services) *Middlewares {
	corsMiddleware := middleware.CorsMiddleware(&middleware.CorsConfig{
		AllowedOrigin: config.CORSOrigin,
	})

	jwtMiddleware := middleware.JwtMiddleware(services.Token)

	errorMiddleware := middleware.ErrorMiddleware()

	rateLimiterMiddleware := middleware.RateLimiterMiddleware(services.RateLimiter, "register", func(r *http.Request) string {
		return r.RemoteAddr
	})

	return &Middlewares{
		CORS:      corsMiddleware,
		JWT:       jwtMiddleware,
		Errors:    errorMiddleware,
		RateLimit: rateLimiterMiddleware,
	}
}
