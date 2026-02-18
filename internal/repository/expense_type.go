package repository

import "time"

type ExpenseType struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
}

type ExpenseTypeRepository interface {
	GetByID(id int) (*ExpenseType, error)
	GetByUserID(id int) ([]*ExpenseType, error)
	Create(t *ExpenseType) error
	DeleteByID(id int) error
}
