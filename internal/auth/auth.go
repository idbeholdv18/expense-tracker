package auth

import (
	"errors"
	"github/idbeholdv18/expense-tracker/internal/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo   repository.UserRepository
	Secret []byte
}

func (s *AuthService) Register(email string, username string, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)

	// TODO: refactor for the only one query to DB
	if err != nil {
		return "", err
	}

	if user != nil {
		return "", errors.New("user already exists")
	}

	user, err = s.Repo.FindByUsername(username)

	if err != nil {
		return "", err
	}

	if user != nil {
		return "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user = &repository.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	err = s.Repo.Create(user)
	if err != nil {
		return "", err
	}

	token, err := CreateToken(user.ID, s.Secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) Login(login string, password string) (string, error) {
	var user *repository.User
	var err error
	if strings.Contains(login, "@") {
		user, err = s.Repo.FindByEmail(login)
	} else {
		user, err = s.Repo.FindByUsername(login)
	}

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
