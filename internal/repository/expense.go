package repository

import "time"

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
	Create(expense *Expense) error
	DeleteByID(userID int, expenseID int) error
	Update(expense *Expense) error
	GetByID(userID int, expenseID int) (*Expense, error)
	GetByUserID(userId int) ([]*Expense, error)
}
