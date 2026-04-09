package httptransport

import "github/idbeholdv18/expense-tracker/internal/contract"

var errorRegistry = make(map[error]func(error) *contract.AppError)

func RegisterError(domainErr error, handler func(error) *contract.AppError) {
	errorRegistry[domainErr] = handler
}
