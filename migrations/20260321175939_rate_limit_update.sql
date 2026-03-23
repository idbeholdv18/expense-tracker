-- +goose Up
DROP TABLE IF EXISTS auth.rate_limit_log;

DROP TYPE IF EXISTS auth.auth_action;

CREATE SCHEMA rate_limit;

CREATE TABLE rate_limit.rate_limit_log (
    id BIGSERIAL PRIMARY KEY,
    key TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX rate_limit_log_key_created_idx ON rate_limit.rate_limit_log (key, created_at);

-- +goose Down
DROP TABLE IF EXISTS rate_limit.rate_limit_log;

DROP SCHEMA IF EXISTS rate_limit;