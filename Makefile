.PHONY: dev_api watch_api fmt

dev_api:
	set -a && . ./.env && set +a && go run ./cmd/api

watch_api:
	reflex -s -r '\.go$$' make dev_api

fmt:
	goimports -w .
	go fmt ./...
