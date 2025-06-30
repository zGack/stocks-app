include .envrc
MIGRATIONS_DIR = ./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -ext sql -dir ${MIGRATIONS_DIR} $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=${MIGRATIONS_DIR} -database=$(DB_MIGRATOR_URL) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=${MIGRATIONS_DIR} -database=$(DB_MIGRATOR_URL) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-force
migrate-force:
	@migrate -path ${MIGRATIONS_DIR} -database $(DB_MIGRATOR_URL) force 1
