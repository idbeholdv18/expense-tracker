package auth

import (
	"encoding/json"
	"github/idbeholdv18/expense-tracker/internal/domain"
	payloadvalidator "github/idbeholdv18/expense-tracker/internal/payload-validator"
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
			Login    string `json:"login" validate:"required"`
			Password string `json:"password" validate:"required"`
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

		setAuthCookie(w, t, 3600)

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
			Username string `json:"username" validate:"required"`
			Email    string `json:"email" validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		errors, err := payloadvalidator.ValidateBody(r, &req)
		if errors != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return json.NewEncoder(w).Encode(&struct {
				Status  int
				Code    string
				Message any
			}{
				Status:  http.StatusBadRequest,
				Code:    "INVALID_REQUEST_BODY",
				Message: errors,
			})
		}
		if err != nil {
			return err
		}

		user, err := h.Auth.Register(r.Context(), req.Email, req.Username, req.Password)
		if err != nil {
			return err
		}

		t, err := h.Token.CreateToken(user.ID)
		if err != nil {
			return err
		}

		setAuthCookie(w, t, 3600)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(map[string]string{
			"token": t,
		})
	}
}

func setAuthCookie(w http.ResponseWriter, token string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   maxAge,
	})
}
