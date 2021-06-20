.PHONY: postgres adminer

postgres:
	sudo docker run --rm -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=gojekIsAwesome postgres

adminer:
	sudo docker run --rm -d -p 8080:8080 adminer