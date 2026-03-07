package user

import (
	"context"
	"database/sql"
	"github/idbeholdv18/expense-tracker/internal/dberrors"
	"github/idbeholdv18/expense-tracker/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, username, password_hash
		FROM users.users
		WHERE email = $1
	`

	user := &User{}

	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
	query := `
		SELECT id, email, username, password_hash
		FROM users.users
		WHERE username = $1
	`

	user := &User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByLogin(ctx context.Context, login string) (*User, error) {
	query := `
		SELECT id, email, username, password_hash
		FROM users.users
		WHERE email=$1 OR username=$2 LIMIT 1
	`

	user := &User{}

	err := r.db.QueryRowContext(ctx, query, login, login).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users.users (email, username, password_hash)
		VALUES($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Username, user.Password).Scan(&user.ID)

	if err != nil {

		if dberrors.IsUniqueViolation(err) {
			return ErrUserAlreadyExists
		}

		if dberrors.IsForeignKeyViolation(err) {
			return domain.ErrInvalidReference
		}

		return err
	}

	return nil
}
