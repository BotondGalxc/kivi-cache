package main

import (
	"context"
	"kivi-cache/cache"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":5001"

type cacheServer struct {
	cache.UnimplementedKiviCacheServiceServer
	values map[string]string
}

func NewCacheServer() *cacheServer {
	server := cacheServer{}
	server.values = make(map[string]string)
	return &server
}

func (server *cacheServer) Get(ctx context.Context, request *cache.GetRequest) (*cache.KeyValue, error) {

	log.Printf("Received request for value %s", request.Key)

	value := server.values[request.Key]

	return &cache.KeyValue{Key: request.Key, Value: value}, nil
}

func (server *cacheServer) Put(ctx context.Context, request *cache.KeyValue) (*cache.PutResponse, error) {
	server.values[request.Key] = request.Value
	return &cache.PutResponse{Result: "Value Stored for Key " + request.Key, Error: ""}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	grpcSrv := grpc.NewServer()
	cache.RegisterKiviCacheServiceServer(grpcSrv, NewCacheServer())
	log.Printf("gRPC server listening at %v", listener.Addr())

	if err := grpcSrv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}

}
