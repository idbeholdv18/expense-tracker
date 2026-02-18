package middleware

import (
	"context"
	"github/idbeholdv18/expense-tracker/internal/auth"
	"net/http"
	"strings"
)

func JwtMiddleware(secret []byte) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := auth.VerifyToken(tokenString, secret)

			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			uidToken, ok := claims["user_id"]
			if !ok {
				http.Error(w, "invalid token: no user_id", http.StatusUnauthorized)
				return
			}

			uidFloat, ok := uidToken.(float64)
			if !ok {
				http.Error(w, "invalid token: user_id wrong type", http.StatusUnauthorized)
				return
			}

			userID := int(uidFloat)

			ctx := context.WithValue(r.Context(), "user_id", userID)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
