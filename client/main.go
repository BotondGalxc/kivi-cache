package main

import (
	"context"
	"kivi-cache/cache"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:5001: %v", err)
	}
	defer conn.Close()

	client := cache.NewKiviCacheServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Get(ctx, &cache.GetRequest{Key: "hello"})
	if err != nil {
		log.Fatalf("error calling function Get: %v", err)
	}

	log.Printf("Response from gRPC server's Get function: %s:%s", response.Key, response.Value)

}
