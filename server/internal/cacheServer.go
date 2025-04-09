package internal

import (
	"context"
	"errors"
	"kivi-cache/cache"
	"log"
)

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
	if request.Key == "" {
		errMessage := "Won't store value without key"
		return &cache.PutResponse{Result: "", Error: errMessage}, errors.New(errMessage)
	}

	if request.Value == "" {
		errMessage := "Won't store empty value."
		return &cache.PutResponse{Result: "", Error: errMessage}, errors.New(errMessage)
	}

	server.values[request.Key] = request.Value
	return &cache.PutResponse{Result: "Value Stored for Key " + request.Key, Error: ""}, nil
}
