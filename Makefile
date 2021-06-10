.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti -p 8080:8080 adminer

migrate:
	migrate -source file://db/migrations \
			-database postgres://postgres:22Andersonst@localhost/postgres?sslmode=disable up

migrate-down:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable down