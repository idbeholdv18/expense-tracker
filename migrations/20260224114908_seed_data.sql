-- +goose Up

INSERT INTO
    users.users (username, email, password)
VALUES (
        'admin',
        'admin@test.com',
        'hashedpassword'
    );

INSERT INTO
    expenses.expense_types (user_id, name)
VALUES (1, 'Food'),
    (1, 'Transport');

-- +goose Down

DELETE FROM expenses.expense_types WHERE user_id = 1;

DELETE FROM users.users WHERE email = 'admin@test.com';