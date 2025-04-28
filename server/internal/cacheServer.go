package internal

import (
	"context"
	"errors"
	"kivi-cache/cache"
	"log/slog"
	"time"
)

const cleanupMiliseconds = 200

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

type cacheServer struct {
	cache.UnimplementedKiviCacheServiceServer
	values     map[string]string
	expiration map[string]time.Time
	log        *slog.Logger
}

func NewCacheServer(logger *slog.Logger) *cacheServer {
	server := cacheServer{log: logger}
	server.values = make(map[string]string)
	server.expiration = make(map[string]time.Time)
	go doEvery(cleanupMiliseconds*time.Millisecond, server.DeleteExpired)
	return &server
}

func NewCacheServerFromMap(items map[string]string, logger *slog.Logger) *cacheServer {
	server := NewCacheServer(logger)
	server.values = items
	return server
}

func (server *cacheServer) Count() int {
	return len(server.values)
}

func (server *cacheServer) Get(ctx context.Context, request *cache.GetRequest) (*cache.KeyValue, error) {

	server.log.Info("Received request for key %s", request.Key, nil)

	value, ok := server.values[request.Key]
	if !ok {
		errMessage := "The key " + request.Key + " does not exist"
		return nil, errors.New(errMessage)
	}

	return &cache.KeyValue{Key: request.Key, Value: value}, nil
}

func (server *cacheServer) Put(ctx context.Context, request *cache.PutRequest) (*cache.PutResponse, error) {
	if request.Key == "" {
		errMessage := "Won't store value without key"
		return nil, errors.New(errMessage)
	}

	if request.Value == "" {
		errMessage := "Won't store empty value."
		return nil, errors.New(errMessage)
	}

	server.values[request.Key] = request.Value
	server.log.Info("Add value for key", "key", request.Key)
	if request.ExpiresSec > 0 {
		server.expiration[request.Key] = time.Now().Add(time.Second * time.Duration(request.ExpiresSec))
	}

	return &cache.PutResponse{Result: "Value Stored for Key " + request.Key, Error: ""}, nil
}

func (server *cacheServer) Delete(ctx context.Context, request *cache.DeleteRequest) (*cache.DeleteResponse, error) {
	delete(server.values, request.Key)
	delete(server.expiration, request.Key)
	server.log.Info("Deleted key", "key", request.Key)
	return &cache.DeleteResponse{Result: "deleted item " + request.Key}, nil
}

func (server *cacheServer) DeleteExpired() {
	for key, exptime := range server.expiration {
		if exptime.Before(time.Now()) {
			delete(server.values, key)
			delete(server.expiration, key)
			server.log.Info("Deleted expired key", "key", key)
		}
	}
}
