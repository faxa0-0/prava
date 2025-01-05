run: build
	@./bin/api

build:
	@go build -o ./bin/api main.go

clean:
	@rm -rf bin

migration:
	@migrate create -ext sql -dir migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run migrate/main.go up

migrate-down:
	@go run migrate/main.go down
