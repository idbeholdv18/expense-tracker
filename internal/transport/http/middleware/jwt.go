package middleware

import (
	"context"
	"github/idbeholdv18/expense-tracker/internal/token"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"net/http"
	"strings"
)

func JwtMiddleware(tokenService *token.TokenService) func(next httptransport.AppHandler) httptransport.AppHandler {
	return func(next httptransport.AppHandler) httptransport.AppHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				return ErrUnauthorized
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := tokenService.VerifyToken(tokenString)

			if err != nil {
				return ErrUnauthorized
			}

			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)

			r = r.WithContext(ctx)

			return next(w, r)
		}
	}

}
