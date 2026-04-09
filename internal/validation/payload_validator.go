package validation

import (
	"github.com/go-playground/validator/v10"
)

type ValidationService struct {
	v *validator.Validate
}

func NewValidationService() *ValidationService {
	v := validator.New()

	return &ValidationService{
		v: v,
	}
}

func (v *ValidationService) Validate(i interface{}) error {
	err := v.v.Struct(i)
	if err == nil {
		return nil
	}

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	fields := map[string]string{}
	for _, fe := range ve {
		fields[fe.Field()] = fe.Tag()
	}

	return NewValidationError(fields)
}
