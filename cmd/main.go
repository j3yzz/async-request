package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Application was connected to Redis successfully.")
	}
}
