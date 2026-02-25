package token

type TokenVerifier interface {
	VerifyToken(tokenString string) (*VerifyTokenDto, error)
}
