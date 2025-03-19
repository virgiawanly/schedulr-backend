SERVICE_NAME = business-service
OUTPUT_BIN = bin/$(SERVICE_NAME)

MIGRATE_DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
MIGRATIONS_DIR=migrations
MIGRATE_CMD=migrate -path $(MIGRATIONS_DIR) -database $(MIGRATE_DB_URL)

build:
	@echo "🔨 Building $(SERVICE_NAME)..."
	@mkdir -p bin
	go build -o $(OUTPUT_BIN) ./cmd/server/
	@echo "✅ Build complete."

run: build
	@echo "🚀 Running $(SERVICE_NAME)..."
	@$(OUTPUT_BIN)

execute:
	@echo "🚀 Running $(SERVICE_NAME)..."
	@$(OUTPUT_BIN)

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf bin/
	@echo "✅ Clean complete."

create-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down 1

migrate-reset:
	$(MIGRATE_CMD) drop -f

migrate-status:
	$(MIGRATE_CMD) version

.PHONY: build run execute clean create-migration migrate-up migrate-down migrate-reset migrate-status