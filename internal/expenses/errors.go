package expenses

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrIncorrectDateFormat           = errors.New("incorrect date format")
	ErrIncorrectExpenseCreatePayload = errors.New("incorrect expense create payload")
	ErrIncorrectExpenseDeletePayload = errors.New("incorrect expense delete payload")
	ErrIncorrectExpenseUpdatePayload = errors.New("incorrect expense update payload")
)

func RegisterErrors(register func(error, func(error) *contract.AppError)) {
	register(ErrIncorrectDateFormat, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"INCORRECT_DATE_FORMAT",
			"incorrect date format",
		)
	})

	register(ErrIncorrectExpenseCreatePayload, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"INCORRECT_PAYLOAD",
			"incorrect payload",
		)
	})

	register(ErrIncorrectExpenseDeletePayload, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"INCORRECT_PAYLOAD",
			"incorrect payload",
		)
	})

	register(ErrIncorrectExpenseUpdatePayload, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"INCORRECT_PAYLOAD",
			"incorrect payload",
		)
	})
}
