include .env

current_dir = $(shell pwd)

start-dev:
	docker compose --profile full up

stop-dev:
	docker compose --profile full down

build-prod:
	docker build -t delivery-order-api:prod .

run-prod:
	docker run --env-file .env  delivery-order-api:prod

local-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

setup-dev:
	migrate -database ${POSTGRESQL_URL_PUBLIC} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL_PUBLIC} -path db/migrations down


test-report:
	go test ./... -v -cover  -coverprofile=c.out && go tool cover -html=c.out

# OWASP ZAP RESPONSE
zap-scan:
	docker run --user root --network host -v ${current_dir}:/zap/wrk/:rw -t zaproxy/zap-stable zap-api-scan.py -t http://localhost:8081/swagger/index.html -f openapi -r reportfix.html  -g zap_config && docker rm zap

zap-scan-products:
	docker run --user root --network host -v ${current_dir}:/zap/wrk/:rw -t zaproxy/zap-stable zap-api-scan.py -t http://localhost:8081/v1/delivery/products -f openapi -r reportfix_products.html  -g zap_config && docker rm zap

zap-scan-orders:
	docker run --user root --network host -v ${current_dir}:/zap/wrk/:rw -t zaproxy/zap-stable zap-api-scan.py -t http://localhost:8081/v1/delivery/orders -f openapi -r reportfix_orders.html  -g zap_config && docker rm zap

