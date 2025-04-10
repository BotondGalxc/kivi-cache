package internal

import (
	"context"
	"kivi-cache/cache"
	"testing"
	"time"
)

func TestCacheServer(t *testing.T) {
	t.Run("put and receive a value", func(t *testing.T) {
		srv := NewCacheServer()
		key := "test"
		value := "123"

		srv.Put(context.Background(), &cache.PutRequest{Key: key, Value: value, ExpiresSec: -1})
		result, _ := srv.Get(context.Background(), &cache.GetRequest{Key: key})
		got := result.Value
		want := value

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("get error on putting empty key", func(t *testing.T) {
		srv := NewCacheServer()
		key := ""
		value := "123"

		_, err := srv.Put(context.Background(), &cache.PutRequest{Key: key, Value: value, ExpiresSec: -1})

		if err == nil {
			t.Errorf("want err not to be nil")
		}
	})

	t.Run("get error on putting empty value", func(t *testing.T) {
		srv := NewCacheServer()
		key := "test"
		value := ""

		_, err := srv.Put(context.Background(), &cache.PutRequest{Key: key, Value: value, ExpiresSec: -1})

		if err == nil {
			t.Errorf("want err not to be nil")
		}
	})

	t.Run("get error on receiving nonexistent key", func(t *testing.T) {
		srv := NewCacheServer()
		key := "test"

		_, err := srv.Get(context.Background(), &cache.GetRequest{Key: key})

		if err == nil {
			t.Errorf("want err not to be nil")
		}
	})

	t.Run("delete item", func(t *testing.T) {
		key := "deleteme"
		srv := NewCacheServerFromMap(map[string]string{key: "obsolete"})

		srv.Delete(context.Background(), &cache.DeleteRequest{Key: key})

		response, err := srv.Get(context.Background(), &cache.GetRequest{Key: key})

		if response != nil {
			t.Errorf("Expected no item, got response with %s:%s", response.Key, response.Value)
		}

		if err == nil {
			t.Errorf("Expected error because of nonexistent item")
		}
	})

	t.Run("items expire", func(t *testing.T) {
		srv := NewCacheServer()
		key := "test"
		value := "123"

		srv.Put(context.Background(), &cache.PutRequest{Key: key, Value: value, ExpiresSec: 1})

		time.Sleep(time.Second * 2)

		response, err := srv.Get(context.Background(), &cache.GetRequest{Key: key})

		if response != nil {
			t.Errorf("Expected no item, got response with %s:%s", response.Key, response.Value)
		}

		if err == nil {
			t.Errorf("Expected error because of nonexistent item")
		}
	})

}
