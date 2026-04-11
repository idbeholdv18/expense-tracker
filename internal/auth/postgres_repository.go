package auth

import (
	"context"
	"database/sql"
	"errors"
	"github/idbeholdv18/expense-tracker/internal/domain"
)

type PostgresRegistrationRequestRepository struct {
	db *sql.DB
}

func NewPostgresRegistrationRequestRepository(db *sql.DB) *PostgresRegistrationRequestRepository {
	return &PostgresRegistrationRequestRepository{
		db: db,
	}
}

func (r *PostgresRegistrationRequestRepository) GetById(ctx context.Context, id int) (*RegistrationRequest, error) {
	query := `
		SELECT id, email, username, token_hash, password_hash, status, expires_at, created_at, updated_at
		FROM auth.registration_request
		WHERE id=$1;
	`

	rr := &RegistrationRequest{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&rr.Id,
		&rr.Email,
		&rr.Username,
		&rr.TokenHash,
		&rr.PasswordHash,
		&rr.Status,
		&rr.ExpiresAt,
		&rr.CreatedAt,
		&rr.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return rr, nil
}

func (r *PostgresRegistrationRequestRepository) UpdateStatusById(ctx context.Context, id int, status RegistrationStatus) error {
	query := `
		UPDATE auth.registration_request
		SET status = $1, updated_at = now()
		WHERE id=$2;
	`

	res, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *PostgresRegistrationRequestRepository) Create(ctx context.Context, rr *RegistrationRequest) error {
	query := `
		INSERT INTO auth.registration_request(email, username, token_hash, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, username, token_hash, password_hash, status, expires_at, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query, rr.Email, rr.Username, rr.TokenHash, rr.PasswordHash).Scan(
		&rr.Id,
		&rr.Email,
		&rr.Username,
		&rr.TokenHash,
		&rr.PasswordHash,
		&rr.Status,
		&rr.ExpiresAt,
		&rr.CreatedAt,
		&rr.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
