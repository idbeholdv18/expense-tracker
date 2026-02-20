package provider

import (
	"encoding/json"
	"errors"
	"github/idbeholdv18/expense-tracker/internal/domain"
	expensetypes "github/idbeholdv18/expense-tracker/internal/expense_types"
	"github/idbeholdv18/expense-tracker/internal/repository"
	"net/http"
)

type ExpenseTypesHandler struct {
	Service *expensetypes.ExpenseTypesService
}

func (h *ExpenseTypesHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int)

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		et := &repository.ExpenseType{
			UserID: userID,
			Name:   req.Name,
		}

		if err := h.Service.Create(r.Context(), et); err != nil {
			switch {
			case errors.Is(err, domain.ErrCategoryAlreadyExists):
				http.Error(w, err.Error(), http.StatusConflict)
			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(et)
	}
}
