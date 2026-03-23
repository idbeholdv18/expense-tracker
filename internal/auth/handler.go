package auth

import (
	"encoding/json"
	"github/idbeholdv18/expense-tracker/internal/domain"
	"github/idbeholdv18/expense-tracker/internal/token"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"net/http"
)

type AuthHandler struct {
	Auth  *AuthService
	Token *token.TokenService
}

func (h *AuthHandler) HandleLogin() httptransport.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodPost {
			return domain.ErrMethodNotAllowed
		}

		var req struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return domain.ErrBadRequest
		}

		user, err := h.Auth.Login(r.Context(), req.Login, req.Password)

		if err != nil {
			return err
		}

		t, err := h.Token.CreateToken(user.ID)

		if err != nil {
			return err
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    t,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   3600,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(map[string]string{
			"token": t,
		})
	}

}

func (h *AuthHandler) HandleRegister() httptransport.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return domain.ErrMethodNotAllowed
		}

		var req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return domain.ErrBadRequest
		}

		user, err := h.Auth.Register(r.Context(), req.Email, req.Username, req.Password)
		if err != nil {
			return err
		}

		t, err := h.Token.CreateToken(user.ID)
		if err != nil {
			return err
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    t,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   3600,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(map[string]string{
			"token": t,
		})
	}

}
