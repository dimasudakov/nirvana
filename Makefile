LOCAL_DB_DSN := "postgres://postgres@localhost:5532/nirvana1?sslmode=disable"

proto:
	protoc -I=./api --go_out=. --go-grpc_out=. nirvana.proto

migrate-up:
	goose -dir ./migrations/postgres postgres $(LOCAL_DB_DSN) up

migrate-down:
	goose -dir ./migrations/postgres postgres $(LOCAL_DB_DSN) down

jet:
	jet -dsn=$(LOCAL_DB_DSN) -path="./internal/app/nirvana/generated/models"
