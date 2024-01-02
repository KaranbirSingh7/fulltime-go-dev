default: run

include Makefile.tools

run-local: install-air
	air

run:
	go run -race main.go

test:
	go test -v ./..
