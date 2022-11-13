include app.env

build:
	docker compose --env-file ./app.env build

up:
	docker compose --env-file ./app.env up

newmig:
	migrate create -ext sql -dir db/migration -seq $(name)
