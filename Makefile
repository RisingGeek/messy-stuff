BINARY_NAME = messy-stuff

build:
	go build -o ${BINARY_NAME} app/main.go
run: build
	./${BINARY_NAME}