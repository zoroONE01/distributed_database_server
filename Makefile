include ./config/.base.env
export $(shell sed 's/=.*//' ./config/.base.env)

server:
	go run cmd/main.go

docker-compose run:
	docker-compose up -d --build
	
docker-compose down:
	docker-compose down