package expenses

import (
	"encoding/json"
	"github/idbeholdv18/expense-tracker/internal/domain"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ExpenseHandler struct {
	Service *ExpenseService
}

func (h *ExpenseHandler) HandleExpense() httptransport.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID, ok := middleware.GetUserID(r.Context())
		if !ok {
			return domain.ErrUnauthorized
		}

		path := strings.TrimPrefix(r.URL.Path, "/api/v1/expenses")

		switch {
		case path == "/" || path == "":
			return h.handleCollection(w, r, userID)
		case strings.HasPrefix(path, "/"):
			return h.handleResource(w, r, userID)
		default:
			return domain.ErrMethodNotAllowed
		}
	}
}

func (h *ExpenseHandler) handleCollection(w http.ResponseWriter, r *http.Request, userID int) error {
	switch r.Method {
	case http.MethodGet:
		return h.handleGetByUserID(w, r, userID)
	case http.MethodPost:
		return h.handleCreate(w, r, userID)
	default:
		return domain.ErrMethodNotAllowed
	}
}

func (h *ExpenseHandler) handleResource(w http.ResponseWriter, r *http.Request, userID int) error {
	switch r.Method {
	case http.MethodDelete:
		return h.handleDelete(w, r, userID)
	case http.MethodPatch:
		return h.handleUpdate(w, r, userID)
	default:
		return domain.ErrMethodNotAllowed
	}
}

func (h *ExpenseHandler) handleGetByUserID(w http.ResponseWriter, r *http.Request, userID int) error {
	expenses, err := h.Service.GetByUserID(r.Context(), userID)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(expenses)
}

func (h *ExpenseHandler) handleCreate(w http.ResponseWriter, r *http.Request, userID int) error {
	var req ExpenseCreateRequestDTO

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return ErrIncorrectExpenseCreatePayload
	}

	expenseDate, err := time.Parse("2006-01-02", req.ExpenseDate)
	if err != nil {
		return ErrIncorrectDateFormat
	}

	payload := &CreateExpenseInput{
		Amount:        req.Amount,
		ExpenseTypeID: req.ExpenseTypeID,
		Currency:      req.Currency,
		Description:   req.Description,
		ExpenseDate:   expenseDate,
	}

	e, err := h.Service.Create(r.Context(), userID, payload)
	if err != nil {
		return err
	}

	res := toCreateResponse(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(res)
}

func (h *ExpenseHandler) handleDelete(w http.ResponseWriter, r *http.Request, userID int) error {
	idString := strings.TrimPrefix(r.URL.Path, "/api/v1/expenses/")

	if strings.Contains(idString, "/") {
		return domain.ErrMethodNotAllowed
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return domain.ErrBadRequest
	}

	if err := h.Service.Delete(r.Context(), userID, id); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *ExpenseHandler) handleUpdate(w http.ResponseWriter, r *http.Request, userID int) error {
	idString := strings.TrimPrefix(r.URL.Path, "/api/v1/expenses/")

	if strings.Contains(idString, "/") {
		return domain.ErrMethodNotAllowed
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return domain.ErrBadRequest
	}

	var req struct {
		Amount        int64  `json:"amount"`
		ExpenseTypeID int    `json:"expense_type_id"`
		Currency      string `json:"currency"`
		Description   string `json:"description"`
		ExpenseDate   string `json:"expense_date"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return ErrIncorrectExpenseUpdatePayload
	}

	expenseDate, err := time.Parse("2006-01-02", req.ExpenseDate)
	if err != nil {
		return ErrIncorrectDateFormat
	}

	expense := &Expense{
		ID:            id,
		UserID:        userID,
		Amount:        req.Amount,
		ExpenseTypeID: req.ExpenseTypeID,
		Currency:      req.Currency,
		Description:   req.Description,
		ExpenseDate:   expenseDate,
	}

	if err := h.Service.Update(r.Context(), userID, expense); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(expense)
}
