DB_URL ?= postgres://idbeholdv:idbeholdv@localhost:5433/expense_tracker?sslmode=disable

GOOSE = go run github.com/pressly/goose/v3/cmd/goose@latest

migrate-up:
	$(GOOSE) -dir migrations postgres "$(DB_URL)" up

migrate-down:
	$(GOOSE) -dir migrations postgres "$(DB_URL)" down

migrate-status:
	$(GOOSE) -dir migrations postgres "$(DB_URL)" status