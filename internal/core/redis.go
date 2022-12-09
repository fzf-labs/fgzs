package core

import "github.com/zeromicro/go-zero/core/stores/redis"

// NewRedis 实例化redis
func NewRedis(conf redis.RedisKeyConf) *redis.Redis {
	return redis.New(conf.Host, func(r *redis.Redis) {
		r.Type = conf.Type
		r.Pass = conf.Pass
	})
}
