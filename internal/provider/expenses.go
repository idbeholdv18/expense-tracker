package provider

import (
	"encoding/json"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	"github/idbeholdv18/expense-tracker/internal/repository"
	"net/http"
	"time"
)

type ExpenseHandler struct {
	Service *expenses.ExpenseService
}

func (h *ExpenseHandler) HandleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int)

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Amount        float64 `json:"amount"`
			ExpenseTypeID int     `json:"expense_type_id"`
			Currency      string  `json:"currency"`
			Description   string  `json:"description"`
			ExpenseDate   string  `json:"expense_date"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		expenseDate, err := time.Parse("2006-01-02", req.ExpenseDate)
		if err != nil {
			http.Error(w, "invalid date format", http.StatusBadRequest)
			return
		}

		expense := &repository.Expense{
			UserID:        userID,
			Amount:        req.Amount,
			ExpenseTypeID: req.ExpenseTypeID,
			Currency:      req.Currency,
			Description:   req.Description,
			ExpenseDate:   expenseDate,
		}

		if err := h.Service.Create(expense); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(expense)
	}
}
