package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	cfg := Config{}
	cfg.context = context.Background()
	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.Redis = initRedis(cfg)
	cfg.wait = &sync.WaitGroup{}

	cfg.serve()
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
