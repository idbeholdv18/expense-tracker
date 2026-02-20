package provider

import (
	"encoding/json"
	"errors"
	"github/idbeholdv18/expense-tracker/internal/domain"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	"github/idbeholdv18/expense-tracker/internal/repository"
	"net/http"
	"time"
)

type ExpenseHandler struct {
	Service *expenses.ExpenseService
}

func (h *ExpenseHandler) HandleExpense() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int)

		switch r.Method {
		case http.MethodPost:
			h.handleCreate(w, r, userID)

		case http.MethodGet:
			h.handleGet(w, r, userID)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

	}
}

func (h *ExpenseHandler) handleGet(w http.ResponseWriter, r *http.Request, userID int) {
	expenses, err := h.Service.GetByUserID(r.Context(), userID)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func (h *ExpenseHandler) handleCreate(w http.ResponseWriter, r *http.Request, userID int) {
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

	if err := h.Service.Create(r.Context(), expense); err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidExpenseType):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}
