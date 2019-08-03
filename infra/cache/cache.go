package cache

import (
	"github.com/go-redis/redis"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// Connect to Redis
func Connect() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Me().Addr(),
		Password: config.Me().Password(),
		DB:       0, // use default DB
	})

	_, err := conn.Ping().Result()
	dont.Panic(err)

	return conn
}
