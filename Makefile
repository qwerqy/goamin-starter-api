include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations
MIGRATE=migrate -path $(MIGRATIONS_PATH) -database "$(DB_ADDR)"

.PHONY: test vet gen-docs migrate-create migrate-up migrate-down migrate-force migrate-status

test:
	@go test -v ./...
	@echo "Tests ran successfully"

vet:
	@go vet ./...
	@echo "Vet ran successfully"

migrate-up:
	@$(MIGRATE) up
	@echo "Migrations applied successfully"

migrate-down:
	@$(MIGRATE) down 1
	@echo "Migrations rolled back successfully"

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Error: Please provide a version number (e.g., make migrate-force version=<version>)"; \
		exit 1; \
	fi
	@$(MIGRATE) force $(version)
	@echo "Migration version forced to $(version)!"

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a migration name (e.g., make migrate-create name=<migration_name>)"; \
		exit 1; \
	fi
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)
	@echo "Migration file created: $(name)"

migrate-status:
	@$(MIGRATE) version
	@echo "Migration status checked!"
# .PHONY: seed
# seed:
# 	@go run cmd/migrate/seed/main.go

gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt
