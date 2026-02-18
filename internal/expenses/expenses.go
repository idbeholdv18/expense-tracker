package expenses

import (
	"github/idbeholdv18/expense-tracker/internal/repository"
)

type ExpenseService struct {
	Repo repository.ExpenseRepository
}

func (s *ExpenseService) Create(
	e *repository.Expense,
) error {
	// TODO: Check if user creates expense for category in his category list
	err := s.Repo.Create(e)

	if err != nil {
		return err
	}

	return nil
}
