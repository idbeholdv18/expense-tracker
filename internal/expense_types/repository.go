package expensetypes

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
	GetByID(ctx context.Context, userID int, id int) (*ExpenseType, error)
	GetByUserID(ctx context.Context, userID int) ([]*ExpenseType, error)
	Create(ctx context.Context, userID int, t *ExpenseType) error
	DeleteByID(ctx context.Context, userID int, id int) error
}
