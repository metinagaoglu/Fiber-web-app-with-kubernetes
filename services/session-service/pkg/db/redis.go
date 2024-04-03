package db

import (
    "github.com/redis/go-redis/v9"
		config "session-service/pkg/config"
)


func ConnectRedis() *redis.Client {
		c, err := config.LoadConfig()
    if err != nil {
        panic(err)
    }

    rdb := redis.NewClient(&redis.Options{
        Addr:     c.RedisUrl,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

		return rdb
}
