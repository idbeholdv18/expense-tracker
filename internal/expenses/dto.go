package expenses

import "time"

type ExpenseCreateRequestDTO struct {
	Amount        int64  `json:"amount"`
	ExpenseTypeID int    `json:"expense_type_id"`
	Currency      string `json:"currency"`
	Description   string `json:"description"`
	ExpenseDate   string `json:"expense_date"`
}

type ExpenseCreateResponseDTO struct {
	ID            int       `json:"id"`
	Amount        int64     `json:"amount"`
	ExpenseTypeID int       `json:"expense_type_id"`
	Currency      string    `json:"currency"`
	Description   string    `json:"description"`
	ExpenseDate   time.Time `json:"expense_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateExpenseInput struct {
	Amount        int64
	ExpenseTypeID int
	Currency      string
	Description   string
	ExpenseDate   time.Time
}

func toCreateResponse(e *Expense) *ExpenseCreateResponseDTO {
	return &ExpenseCreateResponseDTO{
		ID:            e.ID,
		Amount:        e.Amount,
		ExpenseTypeID: e.ExpenseTypeID,
		Currency:      e.Currency,
		Description:   e.Description,
		ExpenseDate:   e.ExpenseDate,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}
