DB_URL=
FILENAME=
VERSION=

run:
	@go run cmd/shopping_cart/main.go
openapi:
	@go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config api/config.yaml ./api/api.yaml
sqlc:
	@sqlc generate --file=db/sqlc.yaml
db_up:
	@docker-compose --env-file app.env up -d
db_down:
	@docker-compose --env-file app.env down -v
migration:
	@migrate create -ext sql -dir db/migrations -seq $(FILENAME)
migrate_up:
	@migrate -path=./db/migrations -database "${DB_URL}" up
migrate_down:
	@migrate -path=./db/migrations -database "${DB_URL}" down
migrate_force:
	@migrate -path=./db/migrations -database "${DB_URL}" force ${VERSION}

.PHONY: run openapi sqlc db_up db_down migration migrate_up migrate_down migrate_force