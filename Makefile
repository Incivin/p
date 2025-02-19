DB_DSN := "postgres://postgres:yourpassword@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations $(name)


migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go