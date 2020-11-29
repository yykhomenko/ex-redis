package main

import (
	"context"
	"testing"

	"github.com/go-redis/redis"
)

func BenchmarkSetGetStd(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}

		if val != "value" {
			b.Fatal("val != 'value'")
		}
	}
}
