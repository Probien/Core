package redis

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	conn *redis.Client
}

func New(opts ...string) *Client {

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     opts[0],
		Password: opts[1],
		DB:       0,
	})

	errHealthCheck := redisConnection.Ping(redisConnection.Context())
	if errHealthCheck.Err() != nil {
		panic(errors.New("Redis connection failed: " + errHealthCheck.Err().Error()))
	}

	return &Client{
		conn: redisConnection,
	}
}

func (c *Client) GetConnection() *redis.Client {
	return c.conn
}
