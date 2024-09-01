package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8/v2"
	"sync"
)

var rdb map[RedisDB]*RedisClient
var once sync.Once

type RedisClient struct {
	*redis.Client
}

func Init(conf *Conf) {
	getOptions := func(db RedisDB) *redis.Options {
		baseOptions := &redis.Options{
			Addr:     conf.Host,
			Username: conf.Username,
			Password: conf.Password,
			DB:       int(db),
		}
		return baseOptions
	}

	once.Do(
		func() {
			rdb = map[RedisDB]*RedisClient{
				RedisDBCache: {
					redis.NewClient(getOptions(RedisDBCache)),
				},
				RedisDBLock: {
					redis.NewClient(getOptions(RedisDBLock)),
				},
			}

			for _, c := range rdb {
				_, err := c.Client.Ping(context.Background()).Result()
				if err != nil {
					panic(err)
				}
				c.Client.AddHook(apmgoredis.NewHook())
			}
		})
}

func GetClient(db RedisDB) *RedisClient {
	if rdb == nil {
		panic("init redis first")
	}
	return rdb[db]
}
