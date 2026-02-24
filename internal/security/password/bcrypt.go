package password

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct {
	Cost int
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		h.Cost,
	)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (h *BcryptHasher) Compare(hashed string, passed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passed))
}
