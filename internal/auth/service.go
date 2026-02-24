package auth

import (
	"context"
	"github/idbeholdv18/expense-tracker/internal/security/password"
	"github/idbeholdv18/expense-tracker/internal/user"
)

type AuthService struct {
	Repo   user.UserRepository
	Hasher password.Hasher
}

func (s *AuthService) Register(ctx context.Context, email string, username string, password string) (*user.User, error) {
	hashedPassword, err := s.Hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	err = s.Repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *AuthService) Login(ctx context.Context, login string, password string) (*user.User, error) {
	u, err := s.Repo.FindByLogin(ctx, login)

	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, ErrInvalidCredentials
	}

	if err := s.Hasher.Compare(u.Password, password); err != nil {
		return nil, ErrInvalidCredentials
	}

	return u, nil
}
