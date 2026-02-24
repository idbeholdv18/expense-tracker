package contract

type HTTPError interface {
	error
	StatusCode() int
	ErrorCode() string
	PublicMessage() string
}
