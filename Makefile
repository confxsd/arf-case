include app.env

build:
	docker compose --env-file ./app.env build

up:
	docker compose --env-file ./app.env up