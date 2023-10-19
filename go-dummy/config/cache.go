package config

import (
	"fmt"
	"go-dummy-project/go-dummy/common"
	"sync"

	"github.com/go-redis/redis"
)

var (
	createRedisConnection sync.Once
	redisClient           *redis.Client
)

func GetRedisClient(channel chan<- struct{}) {
	createRedisConnection.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     common.EnvMap["REDIS_URL"],
			Password: "",
			DB:       0,
		})
	})

	pong, err := redisClient.Ping().Result()
	if err != nil || pong != "PONG" {
		reason := fmt.Sprintf("Error while creating Redis connection : %s", err)
		common.GetLogger().Info(reason)
	}
	common.GetLogger().Info("...Successfully connected to redis docker container...")

	channel <- struct{}{}
}
