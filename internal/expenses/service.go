package expenses

import (
	"context"
	"github/idbeholdv18/expense-tracker/internal/dberrors"
	"github/idbeholdv18/expense-tracker/internal/domain"
)

type ExpenseService struct {
	Repo ExpenseRepository
}

func (s *ExpenseService) Create(
	ctx context.Context,
	userID int,
	e *Expense,
) error {
	e.UserID = userID
	err := s.Repo.Create(ctx, e)

	if err != nil {
		if dberrors.IsForeignKeyViolation(err) {
			return domain.ErrInvalidReference
		}
		return err
	}

	return nil
}

func (s *ExpenseService) GetByUserID(ctx context.Context, userID int) ([]*Expense, error) {
	expenses, err := s.Repo.GetByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (s *ExpenseService) Delete(ctx context.Context, userID int, expenseID int) error {
	return s.Repo.DeleteByID(ctx, userID, expenseID)
}

func (s *ExpenseService) Update(ctx context.Context, userID int, e *Expense) error {
	e.UserID = userID
	return s.Repo.Update(ctx, userID, e)
}

func (s *ExpenseService) GetByID(ctx context.Context, userID int, expenseID int) (*Expense, error) {
	return s.Repo.GetByID(ctx, userID, expenseID)
}
