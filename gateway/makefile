SERVICE_NAME = api-gateway
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

.PHONY: build run execute clean