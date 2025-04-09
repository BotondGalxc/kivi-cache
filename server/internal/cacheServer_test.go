package internal

import (
	"context"
	"kivi-cache/cache"
	"testing"
)

func TestCacheServer(t *testing.T) {
	t.Run("put and receive a value", func(t *testing.T) {
		srv := NewCacheServer()
		key := "test"
		value := "123"

		srv.Put(context.Background(), &cache.KeyValue{Key: key, Value: value})
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

		_, err := srv.Put(context.Background(), &cache.KeyValue{Key: key, Value: value})

		if err == nil {
			t.Errorf("want err not to be nil")
		}
	})

	t.Run("get error on putting empty value", func(t *testing.T) {
		srv := NewCacheServer()
		key := "test"
		value := ""

		_, err := srv.Put(context.Background(), &cache.KeyValue{Key: key, Value: value})

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

}
