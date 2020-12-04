package data

import (
	"github.com/go-redis/redis/v8"
)

func RedisDbConection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "redis_db:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	return rdb
}
