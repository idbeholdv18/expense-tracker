package email

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrEmailWriteError = errors.New("error during write email")
)

func RegisterErrors(register func(error, func(error) *contract.AppError)) {
	register(ErrEmailWriteError, func(error) *contract.AppError {
		return contract.NewAppError(
			http.StatusInternalServerError,
			"EMAIL_WRITING_ERROR",
			"error during write email",
		)
	})
}
