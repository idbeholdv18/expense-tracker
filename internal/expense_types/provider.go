package expensetypes

import (
	"encoding/json"
	"github/idbeholdv18/expense-tracker/internal/domain"
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"net/http"
)

type ExpenseTypesHandler struct {
	Service *ExpenseTypesService
}

func (h *ExpenseTypesHandler) HandleExpenseTypes() httptransport.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID, ok := middleware.GetUserID(r.Context())
		if !ok {
			return domain.ErrUnauthorized
		}

		switch r.Method {
		case http.MethodPost:
			return h.handleCreate(w, r, userID)

		case http.MethodDelete:
			return h.handleDelete(w, r, userID)

		case http.MethodGet:
			return h.handleGetByUserID(w, r, userID)

		default:
			return domain.ErrMethodNotAllowed
		}
	}
}

func (h *ExpenseTypesHandler) handleCreate(w http.ResponseWriter, r *http.Request, userID int) error {
	var req struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return ErrIncorrectExpenseTypeCreatePayload
	}

	et, err := h.Service.Create(r.Context(), userID, req.Name)
	if err != nil {
		return err
	}

	res := &ExpenseTypeCreateResponse{
		ID:        et.ID,
		Name:      et.Name,
		CreatedAt: et.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(res)
}

func (h *ExpenseTypesHandler) handleDelete(w http.ResponseWriter, r *http.Request, userID int) error {
	var req struct {
		ID int `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return ErrIncorrectExpenseTypeCreatePayload
	}

	if err := h.Service.Delete(r.Context(), userID, req.ID); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *ExpenseTypesHandler) handleGetByUserID(w http.ResponseWriter, r *http.Request, userID int) error {
	expenseTypes, err := h.Service.GetByUserID(r.Context(), userID)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(expenseTypes)
}
