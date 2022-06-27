package config

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var Client *redis.Client

func ConnectRedis() {
	err := godotenv.Load("vars.env")

	rsc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if rsc != nil && err != nil {
		log.Fatal(rsc)
		log.Fatal(err)
	}
	Client = rsc
}
