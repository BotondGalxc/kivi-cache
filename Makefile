proto: 
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative ./cache/cache.proto

build-server:
	go build -o bin/kivi-server server/main.go

build-client:
	go build -o bin/kivi-client client/main.go

build: proto build-server build-client