BINARY_NAME=queue-actor

all: build

build:
	mkdir -p out
	go build -o out/${BINARY_NAME} main.go

test:
	go test -v main.go

server:
	./out/${BINARY_NAME} server

worker:
	./out/${BINARY_NAME} worker

clean:
	go clean
	rm ${BINARY_NAME}