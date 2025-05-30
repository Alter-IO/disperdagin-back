up-compose:
	docker compose -p disperdagin -f docker/docker-compose.yaml up -d

down-compose:
	docker compose -p disperdagin -f docker/docker-compose.yaml down

start-service:
	docker compose -p disperdagin -f docker/docker-compose.yaml start

stop-service:
	docker compose -p disperdagin -f docker/docker-compose.yaml stop

migrate-up:
	migrate -path db/migrations -database "postgres://devs:password@localhost:5432/db_disperdagin?sslmode=disable" up

migrate-down:	
	migrate -path db/migrations -database "postgres://devs:password@localhost:5432/db_disperdagin?sslmode=disable" down

migrate-force:	
	migrate -path db/migrations -database "postgres://devs:password@localhost:5432/db_disperdagin?sslmode=disable" force 1

sqlc:
	sqlc generate

logs-api:
	docker logs disperdagin-api -f -n 100

restart:
	docker restart disperdagin-api