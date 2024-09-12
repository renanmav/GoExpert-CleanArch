.PHONY: run test docker-up docker-down proto graphql wire

# Run the application
run:
	go run cmd/main.go cmd/wire_gen.go

# Run tests
test:
	go test ./...

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Generate gRPC code
proto:
	protoc --go_out=. --go-grpc_out=. ./internal/infra/grpc/proto/order.proto

# Generate GraphQL code
graphql:
	go run github.com/99designs/gqlgen generate

# Generate Wire dependency injection code
wire:
	cd cmd && wire
