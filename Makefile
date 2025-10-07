DB_URL=
FILENAME=
VERSION=

# Check if .env file exists and include it
ifneq (,$(wildcard ./app.env))
    include app.env
    export
endif

key_pairs:
	@openssl genpkey -algorithm ed25519 -out private.pem
	@openssl pkey -in private.pem -pubout -out public.pem
run:
	@go run cmd/shopping_cart/main.go
prepare: key_pairs
	@if ! test -f app.env; then \
		cp app.env.example app.env; \
	fi
# make sure to install sqlc locally
sqlc:
	@sqlc generate --file=db/sqlc.yaml
db_up:
	@docker-compose --env-file app.env up -d
db_down:
	@docker-compose --env-file app.env down -v
# make sure to install golang-migrate
migration:
	@migrate create -ext sql -dir db/migrations -seq $(FILENAME)
migrate_up:
	@migrate -path=./db/migrations -database ${DB_URL} up
migrate_down:
	@migrate -path=./db/migrations -database ${DB_URL} down
migrate_force:
	@migrate -path=./db/migrations -database ${DB_URL} force ${VERSION}

.PHONY: run sqlc db_up db_down migration migrate_up migrate_down migrate_force key_pairs prepare
