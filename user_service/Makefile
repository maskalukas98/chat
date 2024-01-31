help:
	@echo "  make migrate-up       : Run database migrations up"
	@echo "  make unit-test        : Run unit tests"
	@echo "  make repository-test  : Run repository tests. Docker-compose must be running."
	@echo "  make api-test  	   : Run api tests. Docker-compose and http server must be running."

migrate-up:
	migrate -path ./migrations -database "postgresql://admin:admin123@localhost:5432/user_db?sslmode=disable" up

unit-test:
	go test ./src/...

repository-test:
	go test ./test/repository

api-test:
	go test ./test/api