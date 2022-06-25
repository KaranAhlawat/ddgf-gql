package repo

import (
	"context"
	"ddgf-new/config"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
)

func NewRedisClient(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf("%s:%s", config.RedisHost, fmt.Sprint(config.RedisPort)),
		Password: config.RedistPassword,
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
