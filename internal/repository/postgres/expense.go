package postgres

import (
	"context"
	"database/sql"
	"github/idbeholdv18/expense-tracker/internal/repository"
)

type PostgresExpensesRepository struct {
	db *sql.DB
}

func NewExpensesRepository(db *sql.DB) *PostgresExpensesRepository {
	return &PostgresExpensesRepository{db: db}
}

func (r *PostgresExpensesRepository) Create(ctx context.Context, e *repository.Expense) error {
	query := `
		INSERT INTO expenses.expenses 
		(user_id, amount, expense_type_id, currency, description, expense_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		e.UserID,
		e.Amount,
		e.ExpenseTypeID,
		e.Currency,
		e.Description,
		e.ExpenseDate,
	).Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt)
}

func (r *PostgresExpensesRepository) DeleteByID(userID int, expenseID int) error {
	_, err := r.db.Exec(`DELETE FROM expenses.expenses WHERE id=$1 AND user_id=$2;`, expenseID, userID)
	return err
}

func (r *PostgresExpensesRepository) Update(e *repository.Expense) error {
	query := `
		UPDATE expenses.expenses
		SET amount=$1,
			expense_type_id=$2,
			currency=$3,
			description=$4,
			expense_date=$5
			udated_at=NOW()
		WHERE id=$6
	`

	_, err := r.db.Exec(
		query,
		e.Amount,
		e.ExpenseTypeID,
		e.Currency,
		e.Description,
		e.ExpenseDate,
		e.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresExpensesRepository) GetByID(userID int, expenseID int) (*repository.Expense, error) {
	query := `
		SELECT id, user_id, amount, expense_type_id, currency, description, expense_date, created_at, updated_at
		FROM expenses.expenses 
		WHERE id=$1 AND user_id=$2
	`

	expense := &repository.Expense{}

	err := r.db.QueryRow(
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
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (r *PostgresExpensesRepository) GetByUserID(ctx context.Context, userID int) ([]*repository.Expense, error) {
	query := `
		SELECT id, user_id, amount, expense_type_id, currency, description, expense_date, created_at, updated_at
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

	expenses := []*repository.Expense{}

	for rows.Next() {
		expense := &repository.Expense{}

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
