package validation

import (
	"errors"
	"fmt"
	"github/idbeholdv18/expense-tracker/internal/contract"
)

var (
	ErrInvalidPayload = errors.New("invalid payload")
)

type ValidationErrorWrapper struct {
	Fields map[string]string
}

func (v ValidationErrorWrapper) Error() string {
	if len(v.Fields) == 0 {
		return "validation failed"
	}

	msg := ""
	for field, err := range v.Fields {
		msg += fmt.Sprintf("%s: %s; ", field, err)
	}
	return msg
}

func RegisterErrors(register func(domainErr error, handler func(error) *contract.AppError)) {
	register(ErrInvalidPayload, func(err error) *contract.AppError {
		ve, ok := err.(ValidationErrorWrapper)
		if !ok {
			return NewValidationError(nil)
		}
		return NewValidationError(ve.Fields)
	})
}

func NewValidationError(errFields map[string]string) *contract.AppError {
	return &contract.AppError{
		Status:  400,
		Code:    "VALIDATION_ERROR",
		Message: "Validation failed",
		Fields:  errFields,
	}
}
