package auth

import (
	"encoding/json"
	"github/idbeholdv18/expense-tracker/internal/domain"
	"github/idbeholdv18/expense-tracker/internal/email"
	"github/idbeholdv18/expense-tracker/internal/token"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"github/idbeholdv18/expense-tracker/internal/validation"
	"net/http"
)

type AuthHandler struct {
	Auth       *AuthService
	Token      *token.TokenService
	Validation *validation.ValidationService
	Email      email.EmailService
}

func (h *AuthHandler) HandleLogin() httptransport.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodPost {
			return domain.ErrMethodNotAllowed
		}

		var payload struct {
			Login    string `json:"login" validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			return domain.ErrBadRequest
		}

		if err := h.Validation.Validate(payload); err != nil {
			return err
		}

		user, err := h.Auth.Login(r.Context(), payload.Login, payload.Password)

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

		var payload struct {
			Username string `json:"username" validate:"required,min=3"`
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,min=8"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			return domain.ErrBadRequest
		}

		if err := h.Validation.Validate(payload); err != nil {
			return err
		}

		user, err := h.Auth.Register(r.Context(), payload.Email, payload.Username, payload.Password)
		if err != nil {
			return err
		}

		t, err := h.Token.CreateToken(user.ID)
		if err != nil {
			return err
		}

		if err := h.Email.Send(r.Context(), email.Email{
			To:      payload.Email,
			Subject: "auth",
			Body:    []byte(t),
		}); err != nil {
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
