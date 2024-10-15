module github.com/gabehamasaki/orders/auth

go 1.22.3

require (
	github.com/gabehamasaki/orders/grpc v0.0.0
	github.com/gabehamasaki/orders/utils v0.0.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.7.1
	github.com/joho/godotenv v1.5.1
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.27.0
	google.golang.org/grpc v1.67.1
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241007155032-5fefd90f89a9 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)

replace github.com/gabehamasaki/orders/grpc v0.0.0 => ../grpc

replace github.com/gabehamasaki/orders/utils v0.0.0 => ../utils
