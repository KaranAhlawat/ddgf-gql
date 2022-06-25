package repo

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalln("redis ping failed.")
	} else {
		log.Println("redis ping succeeded.")
	}

	return rdb
}
