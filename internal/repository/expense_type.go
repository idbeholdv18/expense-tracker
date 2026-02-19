package repository

import (
	"context"
	"time"
)

type ExpenseType struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
}

type ExpenseTypeRepository interface {
	GetByID(ctx context.Context, id int) (*ExpenseType, error)
	GetByUserID(ctx context.Context, id int) ([]*ExpenseType, error)
	Create(ctx context.Context, t *ExpenseType) error
	DeleteByID(ctx context.Context, id int) error
}
