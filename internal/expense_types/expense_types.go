package expensetypes

import (
	"context"
	"github/idbeholdv18/expense-tracker/internal/dberrors"
	"github/idbeholdv18/expense-tracker/internal/domain"
	"github/idbeholdv18/expense-tracker/internal/repository"
)

type ExpenseTypesService struct {
	Repo repository.ExpenseTypeRepository
}

func (s *ExpenseTypesService) Create(ctx context.Context, et *repository.ExpenseType) error {
	err := s.Repo.Create(ctx, et)

	if err != nil {
		if dberrors.IsUniqueViolation(err) {
			return domain.ErrCategoryAlreadyExists
		}
		return err
	}

	return nil
}
