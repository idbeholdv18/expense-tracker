package contract

type AppError struct {
	Status  int
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) StatusCode() int {
	return e.Status
}

func (e *AppError) ErrorCode() string {
	return e.Code
}

func (e *AppError) PublicMessage() string {
	return e.Message
}

func NewAppError(errStatus int, errCode string, errMessage string) *AppError {
	return &AppError{
		Status:  errStatus,
		Code:    errCode,
		Message: errMessage,
	}
}
