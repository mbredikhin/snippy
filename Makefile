watch:
	air -c .air.conf

build:
	docker-compose build app

run:
	docker-compose up app

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go
