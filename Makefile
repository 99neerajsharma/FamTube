GOOS?=linux
GOARCH?=amd64
CGO_ENABLED?=0

SERVER_BINARY_NAME=server
WORKER_BINARY_NAME=worker

build-server:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go build -o build/$(SERVER_BINARY_NAME) cmd/$(SERVER_BINARY_NAME)/main.go
	chmod +x build/$(SERVER_BINARY_NAME)

build-worker:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go build -o build/$(WORKER_BINARY_NAME) cmd/$(WORKER_BINARY_NAME)/main.go
	chmod +x build/$(WORKER_BINARY_NAME)

clean:
	rm -rf build
	rm -rf vendor/*/
	go clean -modcache

configure:
	go mod download
	go mod tidy
	go mod vendor

all:
	$(MAKE) clean
	$(MAKE) configure
	$(MAKE) build-server
	$(MAKE) build-worker

server:
	$(MAKE) clean
	$(MAKE) configure
	$(MAKE) build-server

worker:
	$(MAKE) clean
	$(MAKE) configure
	$(MAKE) build-worker

build-server-run:
	build/$(SERVER_BINARY_NAME) start

build-worker-run:
	build/$(WORKER_BINARY_NAME) start

server-run:
	go run cmd/$(SERVER_BINARY_NAME)/main.go start

worker-run:
	go run cmd/$(WORKER_BINARY_NAME)/main.go start