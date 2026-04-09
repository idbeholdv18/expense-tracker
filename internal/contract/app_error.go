package contract

type AppError struct {
	Status  int               `json:"status"`
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
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
