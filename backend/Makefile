run: build
	@./bin/api

build:
	@go build -o ./bin/api cmd/api/main.go

clean:
	@rm -rf bin

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
