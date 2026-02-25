package expenses

import (
	"context"
	"database/sql"
	"github/idbeholdv18/expense-tracker/internal/domain"
)

type PostgresExpensesRepository struct {
	db *sql.DB
}

func NewExpensesRepository(db *sql.DB) *PostgresExpensesRepository {
	return &PostgresExpensesRepository{db: db}
}

func (r *PostgresExpensesRepository) Create(ctx context.Context, e *Expense) error {
	query := `
		INSERT INTO expenses.expenses 
		(user_id, amount_int, expense_type_id, currency, description, expense_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		e.UserID,
		e.ExpenseTypeID,
		e.Currency,
		e.Description,
		e.ExpenseDate,
	).Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt)
}

func (r *PostgresExpensesRepository) DeleteByID(ctx context.Context, userID int, expenseID int) error {
	res, err := r.db.Exec(`DELETE FROM expenses.expenses WHERE id=$1 AND user_id=$2;`, expenseID, userID)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *PostgresExpensesRepository) Update(ctx context.Context, userID int, e *Expense) error {
	query := `
		UPDATE expenses.expenses
		SET
			amount_int=$1,
			expense_type_id=$2,
			currency=$3,
			description=$4,
			expense_date=$5,
			updated_at=NOW()
		WHERE id=$6 AND user_id=$7
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		e.Amount,
		e.ExpenseTypeID,
		e.Currency,
		e.Description,
		e.ExpenseDate,
		e.ID,
		e.UserID,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *PostgresExpensesRepository) GetByID(ctx context.Context, userID int, expenseID int) (*Expense, error) {
	query := `
		SELECT id, user_id, amount_int, expense_type_id, currency, description, expense_date, created_at, updated_at
		FROM expenses.expenses 
		WHERE id=$1 AND user_id=$2
	`

	expense := &Expense{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		expenseID,
		userID,
	).Scan(
		&expense.ID,
		&expense.UserID,
		&expense.Amount,
		&expense.ExpenseTypeID,
		&expense.Currency,
		&expense.Description,
		&expense.ExpenseDate,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (r *PostgresExpensesRepository) GetByUserID(ctx context.Context, userID int) ([]*Expense, error) {
	query := `
		SELECT id, user_id, amount_int, expense_type_id, currency, description, expense_date, created_at, updated_at
		FROM expenses.expenses 
		WHERE user_id=$1
		ORDER BY expense_date DESC
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*Expense

	for rows.Next() {
		expense := &Expense{}

		if err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.Amount,
			&expense.ExpenseTypeID,
			&expense.Currency,
			&expense.Description,
			&expense.ExpenseDate,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		); err != nil {
			return expenses, err
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}
