.PHONY: build up up-d down restart rebuild logs shell clean ps

COMPOSE=docker compose
MIGRATE_IMAGE=migrate/migrate
MIGRATIONS_DIR=$(PWD)/migrations
GOOSE = go run github.com/pressly/goose/v3/cmd/goose@latest
DB_URL ?= postgres://idbeholdv:idbeholdv@localhost:5433/expense_tracker?sslmode=disable

migrate-up:
	$(GOOSE) -dir migrations postgres "$(DB_URL)" up

migrate-down:
	$(GOOSE) -dir migrations postgres "$(DB_URL)" down

migrate-status:
	$(GOOSE) -dir migrations postgres "$(DB_URL)" status

migrate-create:
	@read -p "Migration name: " name; \
	$(GOOSE) -dir migrations create $$name sql


build:
	$(COMPOSE) build

up:
	$(COMPOSE) up

up-d:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

restart: down up-d

rebuild:
	$(COMPOSE) build --no-cache
	$(COMPOSE) up -d

logs:
	$(COMPOSE) logs -f --tail=100

shell:
	$(COMPOSE) exec authd sh

ps:
	$(COMPOSE) ps

clean:
	$(COMPOSE) down -v --remove-orphans
