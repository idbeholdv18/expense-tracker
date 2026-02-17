package auth

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/dberrors"
	"github/idbeholdv18/expense-tracker/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo   repository.UserRepository
	Secret []byte
}

func (s *AuthService) Register(email string, username string, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &repository.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	err = s.Repo.Create(user)
	if err != nil {
		if dberrors.IsUniqueViolation(err) {
			return "", errors.New("user already exists")
		}
		return "", err
	}

	return CreateToken(user.ID, s.Secret)
}

func (s *AuthService) Login(login string, password string) (string, error) {
	user, err := s.Repo.FindByLogin(login)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("unauthorized")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("unauthorized")
	}

	token, err := CreateToken(user.ID, s.Secret)
	if err != nil {
		return "", err
	}
	return token, nil
}
