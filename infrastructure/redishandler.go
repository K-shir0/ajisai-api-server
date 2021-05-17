package infrastructure

import (
	"github.com/go-redis/redis"
)

func (r *Router) NewRedisHandler() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: r.config.Redis.Host,
	})
}
