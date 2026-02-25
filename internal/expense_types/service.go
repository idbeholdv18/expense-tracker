package expensetypes

import (
	"context"
)

type ExpenseTypesService struct {
	Repo ExpenseTypeRepository
}

func (s *ExpenseTypesService) Create(ctx context.Context, userID int, name string) (*ExpenseType, error) {
	et := &ExpenseType{
		UserID: userID,
		Name:   name,
	}

	if err := s.Repo.Create(ctx, et); err != nil {
		return nil, err
	}

	return et, nil
}

func (s *ExpenseTypesService) Delete(ctx context.Context, userID int, expenseTypeID int) error {
	return s.Repo.DeleteByID(ctx, userID, expenseTypeID)
}

func (s *ExpenseTypesService) Get(ctx context.Context, userID int, expenseTypeID int) (*ExpenseType, error) {
	return s.Repo.GetByID(ctx, userID, expenseTypeID)
}

func (s *ExpenseTypesService) GetByUserID(ctx context.Context, userID int) ([]*ExpenseType, error) {
	return s.Repo.GetByUserID(ctx, userID)
}
