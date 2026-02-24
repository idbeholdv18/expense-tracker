package transport_errors

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/contract"
)

var registry = make(map[error]func(error) *contract.AppError)

func Register(err error, handler func(error) *contract.AppError) {
	registry[err] = handler
}

func Get(err error) (func(error) *contract.AppError, bool) {
	for domainErr, handler := range registry {
		if errors.Is(err, domainErr) {
			return handler, true
		}
	}
	return nil, false
}
