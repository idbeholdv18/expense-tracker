package bootstrap

import (
	"github/idbeholdv18/expense-tracker/internal/config"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
)

type Middlewares struct {
	CORS   middleware.Middleware
	JWT    middleware.Middleware
	Errors middleware.Middleware
}

func RegisterMiddlewares(config *config.Config, services *Services) *Middlewares {
	corsMiddleware := middleware.CorsMiddleware(&middleware.CorsConfig{
		AllowedOrigin: config.CORSOrigin,
	})

	jwtMiddleware := middleware.JwtMiddleware(services.Token)

	errorMiddleware := middleware.ErrorMiddleware()

	return &Middlewares{
		CORS:   corsMiddleware,
		JWT:    jwtMiddleware,
		Errors: errorMiddleware,
	}
}
