package expenses

import (
	"context"
	"github/idbeholdv18/expense-tracker/internal/dberrors"
	"github/idbeholdv18/expense-tracker/internal/domain"
	"github/idbeholdv18/expense-tracker/internal/repository"
)

type ExpenseService struct {
	Repo repository.ExpenseRepository
}

func (s *ExpenseService) Create(
	ctx context.Context,
	e *repository.Expense,
) error {
	err := s.Repo.Create(ctx, e)

	if err != nil {
		if dberrors.IsForeignKeyViolation(err) {
			return domain.ErrInvalidExpenseType
		}
		return err
	}

	return nil
}

func (s *ExpenseService) GetByUserID(ctx context.Context, userID int) ([]*repository.Expense, error) {
	expenses, err := s.Repo.GetByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}
