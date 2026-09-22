.PHONY: dev_api watch_api fmt migrate_up sql_generate

dev_api:
	set -a && . ./.env && set +a && go run ./cmd/api

watch_api:
	reflex -s -r '\.go$$' make dev_api

fmt:
	goimports -w .
	go fmt ./...

migrate_up:
	set -a && . ./.env && set +a && go tool goose -dir db/migrations postgres "$$DATABASE_URL" up

sql_generate:
	go tool sqlc generate