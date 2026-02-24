-- +goose Up

CREATE SCHEMA IF NOT EXISTS expenses;
CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE expenses.expenses (
    id integer NOT NULL,
    user_id integer,
    amount numeric(10, 2) NOT NULL,
    expense_type_id integer,
    currency character varying(3) DEFAULT 'RUB'::character varying NOT NULL,
    description text,
    expense_date date NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE expenses.expense_id_seq AS integer START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE expenses.expense_id_seq OWNED BY expenses.expenses.id;

CREATE TABLE expenses.expense_types (
    id integer NOT NULL,
    user_id integer,
    name character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE expenses.expense_types_id_seq AS integer START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE expenses.expense_types_id_seq OWNED BY expenses.expense_types.id;

CREATE TABLE users.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    password text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE users.users_id_seq AS integer START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE users.users_id_seq OWNED BY users.users.id;

ALTER TABLE ONLY expenses.expense_types
ALTER COLUMN id
SET DEFAULT nextval(
    'expenses.expense_types_id_seq'::regclass
);

ALTER TABLE ONLY expenses.expenses
ALTER COLUMN id
SET DEFAULT nextval(
    'expenses.expense_id_seq'::regclass
);

ALTER TABLE ONLY users.users
ALTER COLUMN id
SET DEFAULT nextval(
    'users.users_id_seq'::regclass
);

ALTER TABLE ONLY expenses.expenses
ADD CONSTRAINT expence_pkey PRIMARY KEY (id);

ALTER TABLE ONLY expenses.expense_types
ADD CONSTRAINT expense_types_pkey PRIMARY KEY (id);

ALTER TABLE ONLY expenses.expense_types
ADD CONSTRAINT expense_types_user_id_id_unique UNIQUE (user_id, id);

ALTER TABLE ONLY expenses.expense_types
ADD CONSTRAINT expense_types_user_id_name_key UNIQUE (user_id, name);

ALTER TABLE ONLY users.users
ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE ONLY users.users
ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY users.users
ADD CONSTRAINT users_username_key UNIQUE (username);

ALTER TABLE ONLY expenses.expense_types
ADD CONSTRAINT expense_types_user_id_fkey FOREIGN KEY (user_id) REFERENCES users.users (id) ON DELETE CASCADE;

ALTER TABLE ONLY expenses.expenses
ADD CONSTRAINT fk_expenses_user FOREIGN KEY (user_id) REFERENCES users.users (id) ON DELETE CASCADE;

ALTER TABLE ONLY expenses.expenses
ADD CONSTRAINT fk_expenses_user_type FOREIGN KEY (user_id, expense_type_id) REFERENCES expenses.expense_types (user_id, id) ON DELETE SET NULL;

-- +goose Down

DROP TABLE IF EXISTS expenses.expenses CASCADE;

DROP TABLE IF EXISTS expenses.expense_types CASCADE;

DROP TABLE IF EXISTS users.users CASCADE;

DROP SEQUENCE IF EXISTS expenses.expense_id_seq CASCADE;

DROP SEQUENCE IF EXISTS expenses.expense_types_id_seq CASCADE;

DROP SEQUENCE IF EXISTS users.users_id_seq CASCADE;