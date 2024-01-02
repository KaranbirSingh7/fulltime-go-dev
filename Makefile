
default: run

run:
	go run -race main.go

test:
	go test -v ./..
