package expenses

import (
	"context"
	"time"
)

type Expense struct {
	ID            int
	UserID        int
	Amount        float64
	AmountInt     int64
	ExpenseTypeID int
	Currency      string
	Description   string
	ExpenseDate   time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ExpenseRepository interface {
	Create(ctx context.Context, expense *Expense) error
	DeleteByID(ctx context.Context, userID int, expenseID int) error
	Update(ctx context.Context, userID int, expense *Expense) error
	GetByID(ctx context.Context, userID int, expenseID int) (*Expense, error)
	GetByUserID(ctx context.Context, userId int) ([]*Expense, error)
}

func (e *Expense) AmountFloat() float64 {
	return float64(e.AmountInt) / 100
}
