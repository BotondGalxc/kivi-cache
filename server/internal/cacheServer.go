package internal

import (
	"context"
	"errors"
	"kivi-cache/cache"
	"log/slog"
	"sync"
	"time"
)

const cleanupMiliseconds = 200

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

type cacheValue struct {
	value      string
	expiration *time.Time
}

type LockableMap struct {
	sync.RWMutex
	m map[string]cacheValue
}

type cacheServer struct {
	cache.UnimplementedKiviCacheServiceServer
	values LockableMap
}

func NewCacheServer() *cacheServer {
	server := cacheServer{}
	server.values = LockableMap{m: make(map[string]cacheValue)}
	go doEvery(cleanupMiliseconds*time.Millisecond, server.DeleteExpired)
	return &server
}

func NewCacheServerFromMap(items map[string]string) *cacheServer {
	server := NewCacheServer()
	for k, v := range items {
		server.values.m[k] = cacheValue{v, nil}
	}
	return server
}

func (server *cacheServer) Count() int {
	return len(server.values.m)
}

func (server *cacheServer) Get(ctx context.Context, request *cache.GetRequest) (*cache.KeyValue, error) {

	slog.Info("Received request for key %s", request.Key, nil)

	server.values.RLock()
	defer server.values.RUnlock()

	value, ok := server.values.m[request.Key]
	if !ok {
		errMessage := "The key " + request.Key + " does not exist"
		return nil, errors.New(errMessage)
	}

	return &cache.KeyValue{Key: request.Key, Value: value.value}, nil
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

	value := cacheValue{request.Value, nil}
	if request.ExpiresSec > 0 {
		time := time.Now().Add(time.Second * time.Duration(request.ExpiresSec))
		value.expiration = &time
	}

	server.values.Lock()
	defer server.values.Unlock()

	server.values.m[request.Key] = value
	slog.Info("Add value for key", "key", request.Key)
	return &cache.PutResponse{Result: "Value Stored for Key " + request.Key, Error: ""}, nil
}

func (server *cacheServer) Delete(ctx context.Context, request *cache.DeleteRequest) (*cache.DeleteResponse, error) {
	server.values.Lock()
	defer server.values.Unlock()

	delete(server.values.m, request.Key)
	slog.Info("Deleted key", "key", request.Key)
	return &cache.DeleteResponse{Result: "deleted item " + request.Key}, nil
}

func (server *cacheServer) DeleteExpired() {
	server.values.Lock()
	defer server.values.Unlock()

	for key, val := range server.values.m {
		if val.expiration != nil {
			if val.expiration.Before(time.Now()) {
				delete(server.values.m, key)
				slog.Info("Deleted expired key", "key", key)
			}
		}
	}
}
