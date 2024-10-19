.PHONY: run build migrate drop-db create-db fresh-db init-db sqlc-gen install-deps buf-lint buf-generate

build:
	@echo "Cleaning up build directories..."
	@rm -rf ./auth/build && mkdir -p ./auth/build
	@rm -rf ./products/build && mkdir -p ./products/build
	@rm -rf ./gateway/build && mkdir -p ./gateway/build
	@echo "Building services..."
	@cd ./auth && cp ./.env ./build/.env && go build -o ./build/auth ./cmd/grpc
	@cd ./products && cp ./.env ./build/.env && go build -o ./build/products ./cmd/grpc
	@cd ./gateway && cp ./.env ./build/.env && go build -o ./build/gateway ./cmd/web
	@echo "Services built successfully!"

run: build
	@clear
	@echo "Starting services concurrently..."
	@concurrently --kill-others --names "auth,product,gateway" \
		"cd ./auth/build && ./auth" \
		"cd ./products/build && ./products" \
		"cd ./gateway/build && ./gateway"

sqlc-gen:
	@echo "Generating SQLC..."
	@cd ./auth && sqlc generate
	@cd ./products && sqlc generate
	@echo "SQLC generated successfully!"

install-deps:
	@echo "Installing dependencies..."
	@cd ./auth && go mod tidy
	@cd ./products && go mod tidy
	@cd ./gateway && go mod tidy
	@echo "Dependencies installed successfully!"
buf-lint:
	@echo "Linting proto files..."
	@cd ./grpc && buf lint
	@echo "Proto files linted successfully!"
buf-generate: buf-lint
	@echo "Generating proto files..."
	@cd ./grpc && buf generate
	@echo "Proto files generated successfully!"

# Database Commands Section
init-db:
	@echo "Starting database..."
	@docker compose up -d postgres
	@make migrate
	@echo "Database started successfully!"
migrate:
	@echo "Migrating database..."
	@tern migrate -c ./auth/internal/db/migrations/tern.conf -m ./auth/internal/db/migrations
	@tern migrate -c ./products/internal/db/migrations/tern.conf -m ./products/internal/db/migrations
	@echo "Database migrated successfully!"
drop-db:
	@echo "Dropping database..."
	@docker compose exec postgres psql -U postgres -c "DROP DATABASE IF EXISTS orders"
	@echo "Database dropped successfully!"
create-db:
	@echo "Creating database..."
	@docker compose exec postgres psql -U postgres -c "CREATE DATABASE orders"
	@echo "Database created successfully!"
fresh-db: drop-db create-db migrate
	@echo "Database refreshed successfully!"
