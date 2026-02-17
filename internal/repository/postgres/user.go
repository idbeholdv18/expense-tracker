package postgres

import (
	"database/sql"
	"github/idbeholdv18/expense-tracker/internal/repository"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindByEmail(email string) (*repository.User, error) {
	query := `
		SELECT id, email, username, password
		FROM users.users
		WHERE email = $1
	`

	user := &repository.User{}

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByUsername(username string) (*repository.User, error) {
	query := `
		SELECT id, email, username, password
		FROM users.users
		WHERE username = $1
	`

	user := &repository.User{}
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByLogin(login string) (*repository.User, error) {
	query := `
		SELECT id, email, username, password
		FROM users.users
		WHERE email=$1 OR username=$2 LIMIT 1
	`

	user := &repository.User{}

	err := r.db.QueryRow(query, login, login).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Create(user *repository.User) error {
	query := `
		INSERT INTO users.users (email, username, password)
		VALUES($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(query, user.Email, user.Username, user.Password).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}
