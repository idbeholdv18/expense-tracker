package repository

type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	// Login is either email or username
	FindByLogin(login string) (*User, error)
	Create(user *User) error
}
