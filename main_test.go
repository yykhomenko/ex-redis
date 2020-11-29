package main

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redis"
	"github.com/mediocregopher/radix"
)

func BenchmarkSetGet(b *testing.B) {
	rad, err := radix.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer rad.Close()

	client := radix.Client(rad)

	key := "k"
	val := "v"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		if err := client.Do(radix.Cmd(nil, "SET", key, val)); err != nil {
			b.Fatal(err)
		}

		var out string
		if err := client.Do(radix.Cmd(&out, "GET", key)); err != nil {
			b.Fatal(err)
		} else if out != val {
			b.Fatal(errors.New("got wrong value"))
		}
	}
}

func BenchmarkSetGetStd(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

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
