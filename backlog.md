# Backlog

## Auth

### Register

1. [backend] validate register payload [o]
2. [backend] create email_verification entry [o]
3. [backend] send verification email [o]
4. [backend] commit user creation on successful verification [o]
5. [backend] resend verification email [o]
6. [backend] expiration handling for verification tokens [o]

### Login

1. [backend] validate login payload [o]
2. [backend] verify password [o]
3. [backend] issue access JWT [o]
4. [backend] issue refresh token (stored hashed in DB) [o]
5. [backend] refresh access token endpoint [o]
6. [backend] logout (invalidate refresh token) [o]

### Restore Password

1. [backend] send restore email [o]
2. [backend] verify restore token [o]
3. [backend] update password_hash [o]
4. [backend] invalidate old refresh tokens [o]

## Accounts

### Create

1. [backend] validate payload [o]
2. [backend] create account (cash | card | savings) [o]
3. [backend] initial balance support [o]

#### account properties

- name [o]
- type (cash | card | savings) [o]
- currency [o]
- description [o]

### List accounts

1. [backend] list user accounts [o]
2. [backend] return current balance [o]

## Transactions

### Create transaction

1. [backend] validate payload [o]
2. [backend] create expense [o]
3. [backend] create income [o]
4. [backend] create transfer between accounts (single DB transaction) [o]
5. [backend] prevent race conditions (SELECT FOR UPDATE / tx) [o]

#### transaction properties

- account_id [o]
- amount [o]
- type (income | expense | transfer) [o]
- category [o]
- description [o]
- occurred_at [o]

### List transactions

1. [backend] filter by date range [o]
2. [backend] filter by category [o]
3. [backend] pagination [o]

## Recurring

### Create recurring transaction

1. [backend] fixed monthly amount [o]
2. [backend] day_of_month [o]
3. [backend] start_date [o]
4. [backend] optional end_date [o]

### Forecast

1. [backend] project balance for current month [o]
2. [backend] project balance for next N months [o]
3. [backend] include recurring transactions in projection [o]
4. [backend] do NOT persist forecast into ledger [o]

## Goals

### Create goal

1. [backend] target_amount [o]
2. [backend] target_date [o]
3. [backend] linked savings account [o]

### Goal progress

1. [backend] calculate current progress [o]
2. [backend] calculate required monthly saving [o]
3. [backend] estimate completion date based on current rate [o]

## Categories

1. [backend] default system categories [o]
2. [backend] user-defined categories [o]
3. [backend] category type (income | expense) [o]
