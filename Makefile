LINKMODE := -extldflags '-static -s -w'

proto: 
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative ./cache/cache.proto

test-server:
	go test kivi-cache/server/internal

build-server:
	CGO_ENABLED=0 go build -o bin/kivi-server \
		-ldflags "$(LINKMODE)" \
		-tags netgo \
		server/main.go

build-client:
	CGO_ENABLED=0 go build -o bin/kivi-client \
		-ldflags "$(LINKMODE)" \
		-tags netgo \
		client/main.go

build: proto test-server build-server build-client