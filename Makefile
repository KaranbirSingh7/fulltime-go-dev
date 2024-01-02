default: run

include Makefile.tools

run-local: install-air docker-compose-up
	air

run:
	go run -race main.go

test:
	go test -v ./..

docker-compose-up: docker-compose-check
	docker compose up -d

docker-compose-down:
	docker compose down
