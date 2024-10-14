.PHONY: dev build

build:
	@echo "Building services..."
	@cd ./auth && go build -o ./build/auth && cp ./.env ./build/.env
	@cd ./gateway && go build -o ./build/gateway && cp ./.env ./build/.env
	@echo "Services built successfully!"

dev: build
	@echo "Starting services concurrently..."
	@concurrently --kill-others --names "auth,gate" \
		"cd ./auth/build && ./auth" \
		"cd ./gateway/build && ./gateway"
