package password

type Hasher interface {
	Hash(password string) (string, error)
	Compare(hashed string, passed string) error
}
