up-compose:
	docker compose -f docker/docker-compose.yaml up -d

down-compose:
	docker compose -f docker/docker-compose.yaml down

start-service:
	docker compose -f docker/docker-compose.yaml start

stop-service:
	docker compose -f docker/docker-compose.yaml stop

migrate-up:
	migrate -path db/migrations -database "postgres://dev_mode:Mypassword123!@localhost:5432/inspektorat?sslmode=disable" up

migrate-down:	
	migrate -path db/migrations -database "postgres://dev_mode:Mypassword123!@localhost:5432/inspektorat?sslmode=disable" down

sqlc:
	sqlc generate

logs-api:
	docker logs alter-io-api -f -n 100