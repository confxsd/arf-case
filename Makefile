include app.env

build:
	docker compose --env-file ./app.env  -p arfcase build

up:
	docker compose --env-file ./app.env up

newmig:
	migrate create -ext sql -dir db/migration -seq $(name)

test:
	go test -v -cover ./...
