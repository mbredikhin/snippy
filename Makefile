watch:
	air -c .air.conf

build:
	docker-compose build

run:
	docker-compose up -d

test:
	go test -v ./...

migrate:
	docker-compose run --rm app migrate -path ./schema -database 'postgres://postgres:qwerty@postgres:5432/postgres?sslmode=disable' up
