package payloadvalidator

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateBody(r *http.Request, dest interface{}) (map[string][]string, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	decoder := json.NewDecoder(bytes.NewReader(bodyBytes))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dest); err != nil {
		return map[string][]string{"body": {err.Error()}}, err
	}

	if err := validate.Struct(dest); err != nil {
		errors := make(map[string][]string)

		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			errors[field] = append(errors[field], tag)
		}

		return errors, err
	}

	return nil, nil
}
