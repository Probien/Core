package config

import (
	"errors"
	"os"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func ConnectRedis() {

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	errHealthCheck := redisConnection.Ping(redisConnection.Context())

	if errHealthCheck.Err() != nil {
		panic(errors.New("Redis connection failed, can't continue: " + redisConnection.String()))
	}

	Client = redisConnection
}
