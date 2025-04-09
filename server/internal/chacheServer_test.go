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
}
