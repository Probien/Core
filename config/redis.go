package config

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func ConnectRedis() {

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	errHealthCheck := redisConnection.Ping(redisConnection.Context())

	if errHealthCheck.Err() != nil {
		panic(errors.New("Redis connection failed, can't continue: " + redisConnection.String()))
	}

	Client = redisConnection
}
