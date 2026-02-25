-- +goose Up

ALTER TABLE expenses.expenses ADD COLUMN amount_int BIGINT;

UPDATE expenses.expenses SET amount_int = (amount * 100)::BIGINT;

ALTER TABLE expenses.expenses ALTER COLUMN amount_int SET NOT NULL;

-- +goose Down

ALTER expenses.expenses DROP COLUMN amount_int;