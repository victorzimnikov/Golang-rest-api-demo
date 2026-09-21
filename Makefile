.PHONY: dev_api watch_api

dev_api:
	set -a && source .env && set +a &&go run ./cmd/api

watch_api:
	reflex -s -r '\.go$$' make dev_api