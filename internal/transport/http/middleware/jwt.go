package middleware

import (
	"github/idbeholdv18/expense-tracker/internal/token"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"net/http"
	"strings"
)

func JwtMiddleware(verifier token.TokenVerifier) func(next httptransport.AppHandler) httptransport.AppHandler {
	return func(next httptransport.AppHandler) httptransport.AppHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			authHeader := r.Header.Get("Authorization")

			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 || parts[0] != "Bearer" {
				return ErrUnauthorized
			}

			tokenString := parts[1]

			if tokenString == "" {
				return ErrUnauthorized
			}

			claims, err := verifier.VerifyToken(tokenString)

			if err != nil {
				return ErrUnauthorized
			}

			ctx := SetUserID(r.Context(), claims.UserID)

			return next(w, r.WithContext(ctx))
		}
	}

}
