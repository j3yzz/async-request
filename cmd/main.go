package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	cfg := Config{}
	cfg.context = context.Background()
	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.Redis = initRedis(cfg)
	cfg.wait = &sync.WaitGroup{}
	cfg.waitLock = &sync.Mutex{}
	setDefaultKeyAndValueForRedis(cfg)

	cfg.serve()
}

func setDefaultKeyAndValueForRedis(cfg Config) {
	val := rand.Intn(30 - 1)
	cfg.Redis.Set(cfg.context, "request_counter", val, time.Second*3000)
	cfg.InfoLog.Println("request_counter default value:", val)
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func initRedis(cfg Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(cfg.context).Result()
	if err != nil {
		cfg.ErrorLog.Fatal(err)
	} else {
		cfg.InfoLog.Println("Application was connected to Redis successfully.")
	}

	return rdb
}
