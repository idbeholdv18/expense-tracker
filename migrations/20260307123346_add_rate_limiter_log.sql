-- +goose Up

CREATE TYPE auth.auth_action AS ENUM (
    'login',
    'register',
    'resend',
    'verify',
    'reset_password'
);

CREATE TABLE auth.rate_limit_log (
    id BIGSERIAL PRIMARY KEY,
    key TEXT NOT NULL,
    action auth.auth_action NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX rate_limit_log_key_action_created_idx ON auth.rate_limit_log (key, action, created_at);

-- +goose Down
DROP INDEX IF EXISTS rate_limit_log_key_action_created_idx;

DROP TYPE IF EXISTS auth.auth_action;

DROP TABLE IF EXISTS auth.rate_limit_log;