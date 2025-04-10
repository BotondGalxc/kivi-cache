package internal

import (
	"context"
	"kivi-cache/cache"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type cacheClient struct {
	host               string
	port               string
	credentials        credentials.TransportCredentials
	cacheServiceClient cache.KiviCacheServiceClient
}

type KeyValue struct {
	Key   string
	Value string
}

func NewClient(host string, port string, credentials credentials.TransportCredentials) (*cacheClient, error) {
	client := cacheClient{host: host, port: port, credentials: credentials}

	conn, err := grpc.NewClient(host+":"+port, grpc.WithTransportCredentials(credentials))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at %s:%s: %v", host, port, err)
	}

	client.cacheServiceClient = cache.NewKiviCacheServiceClient(conn)

	return &client, nil
}

func (client *cacheClient) Get(key string) KeyValue {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.cacheServiceClient.Get(ctx, &cache.GetRequest{Key: key})
	if err != nil {
		log.Fatalf("Error on Get: %v", err)
	}

	return KeyValue{response.Key, response.Value}
}
