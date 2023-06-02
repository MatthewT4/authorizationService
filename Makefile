run:
	go run cmd/main.go

migration_up:
	migrate -path internal/database/migration -database "postgresql://postgres:1234@localhost:5432/authorizationSystemTEST?sslmode=disable" -verbose up

migration_down:
	migrate -path internal/database/migration -database "postgresql://postgres:1234@localhost:5432/authorizationSystemTEST?sslmode=disable" -verbose down

migration_force:
	migrate -path internal/database/migration -database "postgresql://postgres:1234@localhost:5432/authorizationSystemTEST?sslmode=disable" -verbose force 1