package domain

import "errors"

var (
	ErrCategoryAlreadyExists = errors.New("category already exists")
	ErrInvalidCategoryName   = errors.New("invalid category name")
	ErrInvalidExpenseType    = errors.New("invalid expense type")
)
