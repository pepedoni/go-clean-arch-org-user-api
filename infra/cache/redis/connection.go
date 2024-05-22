package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

const (
	REDIS_HOST = "REDIS_HOST"
	REDIS_PORT = "REDIS_PORT"
)

var (
	redis_host = os.Getenv(REDIS_HOST)
	redis_port = os.Getenv(REDIS_PORT)
	redis_uri  = fmt.Sprintf("redis://%s:%s/0", redis_host, redis_port)
)

func GetConnection() *redis.Client {
	ctx := context.Background()
	opt, err := redis.ParseURL(redis_uri)
	if err != nil {
		log.Fatalln("Redis connection was refused")
		panic(err)
	}
	rdb := redis.NewClient(opt)

	{
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			log.Fatalln("Redis connection was refused")
			panic(err)
		}
	}

	return rdb
}
