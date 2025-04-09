package main

import (
	"kivi-cache/cache"
	"kivi-cache/server/internal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":5001"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	grpcSrv := grpc.NewServer()
	cache.RegisterKiviCacheServiceServer(grpcSrv, internal.NewCacheServer())
	log.Printf("gRPC server listening at %v", listener.Addr())

	if err := grpcSrv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}

}
