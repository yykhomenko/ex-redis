package main

import (
	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()
}
