run:
	go run main.go

test:
	go test -cover ./...
	
tidy:
	go mod tidy
	
doc:
	echo "Starting swagger generating"
	swag fmt
	swag init -g main.go --pd
	
migrate:
	go run migrations/migrate.go

compose-up:
	docker-compose -f ./builders/docker-compose.yml up -d --build

compose-down:
	docker-compose -f ./builders/docker-compose.yml down
