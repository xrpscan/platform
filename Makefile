SERVICE_NAME=platform
BIN_DIR=bin
MKDIR=mkdir -p

all: build

build:
	${MKDIR} ${BIN_DIR}
	go build -o ${BIN_DIR}/${SERVICE_NAME}-server main.go
	go build -o ${BIN_DIR}/${SERVICE_NAME}-cli cmd/cli/main.go

clean:
	go clean
	rm -rf ${BIN_DIR}
