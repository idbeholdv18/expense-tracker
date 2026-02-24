package user

import (
	"context"
)

type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	// Login is either email or username
	FindByLogin(ctx context.Context, login string) (*User, error)
	Create(ctx context.Context, user *User) error
}
