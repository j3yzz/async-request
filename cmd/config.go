package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

type Config struct {
	Redis    *redis.Client
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	wait     *sync.WaitGroup
	context  context.Context
}
