package expensetypes

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrExpenseTypeAlreadyExists          = errors.New("expense type already exists")
	ErrIncorrectExpenseTypeCreatePayload = errors.New("incorrect expense_type create payload")
)

func RegisterErrors(register func(error, func(error) *contract.AppError)) {

	register(ErrExpenseTypeAlreadyExists, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusConflict,
			"EXPENSE_TYPE_ALREADY_EXISTS",
			"expense type already exists",
		)
	})

	register(ErrIncorrectExpenseTypeCreatePayload, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"INCORRECT_PAYLOAD",
			"incorrect payload",
		)
	})
}
