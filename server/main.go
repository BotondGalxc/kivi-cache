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
}

func (server *cacheServer) Get(ctx context.Context, in *cache.GetRequest) (*cache.KeyValue, error) {
	log.Printf("Received request for value %s", in.Key)
	return &cache.KeyValue{Key: in.Key, Value: "Hello World!"}, nil
}

func (server *cacheServer) Put(ctx context.Context, in *cache.KeyValue) (*cache.PutResponse, error) {
	return &cache.PutResponse{Result: "all good!", Error: ""}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	grpcSrv := grpc.NewServer()
	cache.RegisterKiviCacheServiceServer(grpcSrv, &cacheServer{})
	log.Printf("gRPC server listening at %v", listener.Addr())

	if err := grpcSrv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}

}
