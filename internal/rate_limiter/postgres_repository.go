package ratelimiter

import (
	"context"
	"database/sql"
	"time"
)

type RateLimitLogPostgresRepository struct {
	db *sql.DB
}

func NewLimitLogRepository(db *sql.DB) *RateLimitLogPostgresRepository {
	return &RateLimitLogPostgresRepository{
		db: db,
	}
}

func (r *RateLimitLogPostgresRepository) Count(ctx context.Context, key string, action string, window time.Duration) (int, error) {
	var count int

	query := `
		SELECT COUNT(*)
		FROM auth.rate_limit_log
		WHERE key=$1 AND action=$2 AND created_at > now() - $3::interval
	`

	err := r.db.QueryRowContext(ctx, query, key, action, window.String()).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *RateLimitLogPostgresRepository) Create(ctx context.Context, rll *RateLimitLog) error {
	query := `
		INSERT INTO auth.rate_limit_log (key, action)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	err := r.db.QueryRowContext(ctx, query, rll.Key, rll.Action).Scan(&rll.ID, &rll.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
