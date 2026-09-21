.PHONY: dev_api

dev_api:
	go run ./cmd/api

watch_api:
	reflex -s -r '\.go$$' make dev_api