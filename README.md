run migrator:
go run ./cmd/migrator/ --storage-path=./storage/sso.db --migrations-path=./migrations

run app:
go run cmd/sso/main.go --config=./config/local.yaml
