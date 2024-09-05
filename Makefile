BINARY_NAME=scm
BINARIES_PATH=./bin

all: build-mac build-linux build-windows

build:
	go build -o ${BINARIES_PATH}/${BINARY_NAME} ./cmd/scm/main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o ${BINARIES_PATH}/${BINARY_NAME}-mac ./cmd/scm/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARIES_PATH}/${BINARY_NAME}-linux ./cmd/scm/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${BINARIES_PATH}/${BINARY_NAME}-windows.exe ./cmd/scm/main.go

run: build
	./${BINARIES_PATH}/${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARIES_PATH}/${BINARY_NAME}
