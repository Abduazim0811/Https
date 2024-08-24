run:
	go run cmd/main.go

swag:
	swag init -g internal/https/api/handler/handler.go