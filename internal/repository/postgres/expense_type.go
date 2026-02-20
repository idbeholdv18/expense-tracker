package postgres

import (
	"context"
	"database/sql"
	"errors"
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

func (r *PostgresExpenseTypeRepository) DeleteByID(ctx context.Context, userID int, expenseTypeID int) error {
	query := `
		DELETE FROM expenses.expense_types
		WHERE user_id=$1 AND id=$2;
	`

	res, err := r.db.ExecContext(ctx, query, userID, expenseTypeID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("expense type not found or not owned by user")
	}

	return nil
}

func (r *PostgresExpenseTypeRepository) GetByID(ctx context.Context, userID int, expenseTypeID int) (*repository.ExpenseType, error) {
	query := `
		SELECT id, user_id, name, created_at
		FROM expenses.expense_types
		WHERE user_id=$1 AND id=$2;
	`

	expenseType := &repository.ExpenseType{}

	err := r.db.
		QueryRowContext(ctx, query, userID, expenseTypeID).
		Scan(&expenseType.ID, &expenseType.UserID, &expenseType.Name, &expenseType.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return expenseType, nil
}

func (r *PostgresExpenseTypeRepository) GetByUserID(ctx context.Context, userID int) ([]*repository.ExpenseType, error) {
	query := `
		SELECT id, user_id, name, created_at
		FROM expenses.expense_types
		WHERE user_id=$1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenseTypes []*repository.ExpenseType

	for rows.Next() {
		expenseType := &repository.ExpenseType{}
		err := rows.Scan(&expenseType.ID, &expenseType.UserID, &expenseType.Name, &expenseType.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenseTypes = append(expenseTypes, expenseType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenseTypes, nil
}
