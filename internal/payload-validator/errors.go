package payloadvalidator

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
	"net/http"
)

var (
	ErrInvalidPayload = errors.New("Invalid payload")
)

func RegisterErrors(register func(domainErr error, handler func(error) *contract.AppError)) {
	register(ErrInvalidPayload, func(err error) *contract.AppError {
		return contract.NewAppError(
			http.StatusBadRequest,
			"INVALID PAYLOAD",
			"Invalid payload",
		)
	})
}
