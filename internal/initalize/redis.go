package initalize

import (
	"context"
	"log"
	"my_shop/global"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func InitRedis() {
	redisConfig := global.Config.RedisCache
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
    Password: redisConfig.Password, // no password set
    DB:       redisConfig.DB,  // use default DB
	})
}

// CheckRedisConnection pings Redis to check the connection status
func CheckRedisConnection() bool {
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis connection error: %v", err)
		return false
	}
	log.Println("Redis connected successfully")
	return true
}