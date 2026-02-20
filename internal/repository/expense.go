package repository

import (
	"context"
	"time"
)

type Expense struct {
	ID            int
	UserID        int
	Amount        float64
	ExpenseTypeID int
	Currency      string
	Description   string
	ExpenseDate   time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ExpenseRepository interface {
	Create(ctx context.Context, expense *Expense) error
	DeleteByID(userID int, expenseID int) error
	Update(expense *Expense) error
	GetByID(userID int, expenseID int) (*Expense, error)
	GetByUserID(ctx context.Context, userId int) ([]*Expense, error)
}
