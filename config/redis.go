package config

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	conn *redis.Client
}

func NewRedisClient(opts ...string) *RedisClient {

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     opts[0],
		Password: opts[1],
		DB:       0,
	})

	errHealthCheck := redisConnection.Ping(redisConnection.Context())
	if errHealthCheck.Err() != nil {
		panic(errors.New("Redis connection failed: " + errHealthCheck.Err().Error()))
	}

	return &RedisClient{
		conn: redisConnection,
	}
}

func (r *RedisClient) GetConnection() *redis.Client {
	return r.conn
}
