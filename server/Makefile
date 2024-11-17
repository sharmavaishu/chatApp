postgresinit:
     docker run --name postgre15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgre15 psql

createdb:
	docker exec -it postgre15 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgre15 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down
.PHONY: postgresinit