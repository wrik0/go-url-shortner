postgres:
	docker run --name go-url-shortner-db -p 5432:5432 -e POSTGRES_PASSWORD=SECRET \
	-e POSTGRES_USER=go -e POSTGRES_DB=url-shortner -d postgres:alpine
createdb:
	docker exec -it go-url-shortner-db createdb --username=go-tut --owner=go-tut simple-bank
dropdb:
	docker exec -it go-url-shortner-db dropdb simple-bank
migrateup:
	migrate -path db/migration -database "postgresql://go:SECRET@localhost:5432/url-shortner?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://go:SECRET@localhost:5432/url-shortner?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test