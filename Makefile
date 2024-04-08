watch:
	air -c .air.conf

build:
	docker compose build

run:
	docker compose up

test:
	go test -v ./...

migrate-up:
	docker compose run --rm app migrate -path ./schema -database 'postgres://postgres:qwerty@postgres:5432/postgres?sslmode=disable' up

migrate-down:
	docker compose run --rm app migrate -path ./schema -database 'postgres://postgres:qwerty@postgres:5432/postgres?sslmode=disable' down

db-seed:
	docker compose exec -it db psql -U postgres -d postgres -a -f /schema/seed.sql
