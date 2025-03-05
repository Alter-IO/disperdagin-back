up-compose:
	docker compose -f docker/docker-compose.yaml up -d

down-compose:
	docker compose -f docker/docker-compose.yaml down

start-service:
	docker compose -f docker/docker-compose.yaml start

stop-service:
	docker compose -f docker/docker-compose.yaml stop

migrate-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

migrate-down:	
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

migrate-force:	
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" force 1

sqlc:
	sqlc generate

logs-api:
	docker logs alter-io-api -f -n 100