.PHONY: run build migrate drop-db create-db fresh-db init-db sqlc-gen

build:
	@echo "Building services..."
	@cd ./auth && go build -o ./build/auth && cp ./.env ./build/.env
	@cd ./products && go build -o ./build/products && cp ./.env ./build/.env
	@cd ./gateway && go build -o ./build/gateway && cp ./.env ./build/.env
	@echo "Services built successfully!"

run: build
	@echo "Starting services concurrently..."
	@concurrently --kill-others --names "auth,products,gate" \
		"cd ./auth/build && ./auth" \
		"cd ./products/build && ./products" \
		"cd ./gateway/build && ./gateway"

sqlc-gen:
	@echo "Generating SQLC..."
	@cd ./auth && sqlc generate
	@cd ./products && sqlc generate
	@echo "SQLC generated successfully!"

# Database Commands Section
init-db:
	@echo "Starting database..."
	@docker compose up -d postgres
	@make migrate
	@echo "Database started successfully!"
migrate:
	@echo "Migrating database..."
	@tern migrate -c ./auth/db/migrations/tern.conf -m ./auth/db/migrations
	@tern migrate -c ./products/db/migrations/tern.conf -m ./products/db/migrations
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
