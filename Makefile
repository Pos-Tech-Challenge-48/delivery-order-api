include .env



start-dev:
	docker compose --profile full up

stop-dev:
	docker compose --profile full down

build-prod:
	docker build -t delivery-api:prod .

run-prod:
	docker run --env-file .env  delivery-api:prod

local-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

setup-dev:
	migrate -database ${POSTGRESQL_URL_PUBLIC} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL_PUBLIC} -path db/migrations down
