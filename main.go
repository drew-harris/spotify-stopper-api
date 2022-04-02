package main

import (
	"context"
	"fmt"
	handler "github.com/drew-harris/spotify-stopper-api/api"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var ctx = context.Background()

func main() {
	// Set up redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-14793.c124.us-central1-1.gce.cloud.redislabs.com:14793",
		Password: "Pr2soSboKbJbbQv8KHwf2urwJyuk5Yyh",
		DB:       0,
	})

	err := rdb.Set(ctx, "test", "test", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
	// Set up http
	http.HandleFunc("/", handler.Test)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
