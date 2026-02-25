-- +goose Up
ALTER TABLE expenses.expenses DROP COLUMN amount;

ALTER TABLE expenses.expenses RENAME COLUMN amount_int TO amount;

-- +goose Down
ALTER TABLE expenses.expenses ADD COLUMN amount_int BIGINT;

UPDATE expenses.expenses SET amount_int = amount

ALTER TABLE expenses.expenses
ALTER COLUMN amount TYPE NUMERIC(10, 2) USING (amount_int / 100.0);