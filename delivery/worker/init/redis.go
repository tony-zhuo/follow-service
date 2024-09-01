package init

import (
	"github.com/tony-zhuo/follow-service/pkg/config"
	"github.com/tony-zhuo/follow-service/pkg/redis"
)

func initRedisConf() *redis.Conf {
	path := "./conf"
	name := "redis"
	configType := "yml"
	conf := &redis.Conf{}
	config.Load(path, name, configType, conf)
	return conf
}
