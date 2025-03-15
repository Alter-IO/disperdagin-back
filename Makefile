up-compose:
	docker compose -f docker/docker-compose.yaml up -d

down-compose:
	docker compose -f docker/docker-compose.yaml down

start-service:
	docker compose -f docker/docker-compose.yaml start

stop-service:
	docker compose -f docker/docker-compose.yaml stop

migrate-up:
	migrate -path db/migrations -database "postgres://devs:Mypassword123!@localhost:5432/db_disperdagin?sslmode=disable" up

migrate-down:	
	migrate -path db/migrations -database "postgres://devs:Mypassword123!@localhost:5432/db_disperdagin?sslmode=disable" down

migrate-force:	
	migrate -path db/migrations -database "postgres://devs:Mypassword123!@localhost:5432/db_disperdagin?sslmode=disable" force 1

sqlc:
	sqlc generate

logs-api:
	docker logs disperdagin-api -f -n 100