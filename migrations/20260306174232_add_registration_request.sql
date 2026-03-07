-- +goose Up
CREATE SCHEMA IF NOT EXISTS auth;

CREATE TYPE auth.registration_status AS ENUM (
    'pending',
    'confirmed',
    'expired',
    'cancelled'
);

CREATE TABLE auth.registration_request (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(50) NOT NULL,
    token_hash TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    status auth.registration_status NOT NULL DEFAULT 'pending',
    expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);

ALTER TABLE users.users
ALTER COLUMN email TYPE VARCHAR(255) USING email::VARCHAR(255);

ALTER TABLE users.users RENAME COLUMN password TO password_hash;

CREATE UNIQUE INDEX ON auth.registration_request (email)
WHERE
    status = 'pending';

CREATE UNIQUE INDEX ON auth.registration_request (token_hash);

-- +goose Down
DROP TABLE IF EXISTS auth.registration_request;

DROP TYPE IF EXISTS auth.registration_status;

DROP SCHEMA IF EXISTS auth;