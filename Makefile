.PHONY: postgres adminer kill-all list-all

postgres:
	sudo docker run --rm -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=12345 postgres

adminer:
	sudo docker run --rm -d -ti --network host --name my-adminer adminer

kill-all: 
	sudo docker stop my-postgres && sudo docker stop my-adminer

list-all:
	sudo docker ps -a