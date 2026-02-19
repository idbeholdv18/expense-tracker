package postgres

import (
	"context"
	"database/sql"
	"github/idbeholdv18/expense-tracker/internal/repository"
)

type PostgresExpenseTypeRepository struct {
	db *sql.DB
}

func (r *PostgresExpenseTypeRepository) Create(ctx context.Context, t *repository.ExpenseType) error {
	query := `
		INSERT INTO expenses.expense_types (user_id, name)
		VALUES ($1, $2)
		RETURNING id, created_at;
	`

	return r.db.QueryRowContext(ctx, query, t.UserID, t.Name).Scan(&t.ID, &t.CreatedAt)
}
