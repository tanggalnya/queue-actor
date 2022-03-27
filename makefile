BINARY_NAME=main.out

all: build

build:
	mkdir -p out
	go build -o out/${BINARY_NAME} main.go

test:
	go test -v main.go

server:
	mkdir -p out
	go build -o out/${BINARY_NAME} main.go
	./out/${BINARY_NAME} server

worker:
	mkdir -p out
	go build -o out/${BINARY_NAME} main.go
	./out/${BINARY_NAME} worker

clean:
	go clean
	rm ${BINARY_NAME}