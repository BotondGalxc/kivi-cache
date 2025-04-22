LINKMODE := -linkmode external -extldflags '-static -s -w'

proto: 
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative ./cache/cache.proto

test-server:
	go test kivi-cache/server/internal

build-server:
	go build -o bin/kivi-server \
		-ldflags "$(LINKMODE)" \
		-tags netgo \
		server/main.go

build-client:
	go build -o bin/kivi-client \
		-ldflags "$(LINKMODE)" \
		-tags netgo \
		client/main.go

build: proto test-server build-server build-client