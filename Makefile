.PHONY: dev_api watch_api fmt migrate_up sql_generate check swagger_gen

dev_api:
	set -a && . ./.env && set +a && go run ./cmd/api

watch_api:
	reflex -s -r '\.go$$' make dev_api

fmt:
	goimports -w .
	go fmt ./...

check:
	go vet ./... && go test ./...

migrate_up:
	set -a && . ./.env && set +a && go tool goose -dir db/migrations postgres "$$DATABASE_URL" up

sql_generate:
	go tool sqlc generate

swagger_fmt:
	go tool swag fmt

swagger_gen:
	go tool swag init -g cmd/api/main.go -o docs --parseInternal --useStructName