package infrastructure

import (
	"github.com/go-redis/redis/v8"
)

func (r *Router) NewRedisHandler() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.config.Redis.Host + ":" + r.config.Redis.Port,
		Password: r.config.Redis.Password,
		DB:       0,
	})
}
