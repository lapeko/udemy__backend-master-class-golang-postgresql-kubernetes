MIGRATIONS_PATH=storage/migrations
GO_PATH=/mnt/e/go
DB_CONTAINER_NAME=postgres_alpine
DB_USER=root
DB_PASSWORD=1234
DB_PORT=5432
DB_NAME=simple_bank
DB_CONNECTION_URL=postgres://root:1234@127.0.0.1:5432/root?sslmode=disable

.PHONY:
	migrate docker-create docker-drop docker-start docker-stop db-create dp-drop migrate-gen migrate-up migrate-down

docker-create:
	docker run --name $(DB_CONTAINER_NAME) -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:alpine

docker-drop:
	docker rm $(DB_CONTAINER_NAME)

docker-start:
	docker start $(DB_CONTAINER_NAME)

docker-stop:
	docker stop $(DB_CONTAINER_NAME)

db-create:
	docker exec -i $(DB_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dp-drop:
	docker exec -i $(DB_CONTAINER_NAME) dropdb --username=$(DB_USER) $(DB_NAME)

migrate-gen:
	$(GO_PATH)/bin/migrate.exe create -ext sql -dir $(MIGRATIONS_PATH) -seq init_schema

migrate-up:
	${GO_PATH}/bin/migrate.exe -path $(MIGRATIONS_PATH) -database $(DB_CONNECTION_URL) up

migrate-down:
	${GO_PATH}/bin/migrate.exe -path $(MIGRATIONS_PATH) -database $(DB_CONNECTION_URL) down
