package init

import (
	"github.com/tony-zhuo/follow-service/pkg/db"
	"github.com/tony-zhuo/follow-service/pkg/redis"
	"sync"
)

var conf *Conf
var confOnce sync.Once

type Conf struct {
	App   *App
	DB    *db.DBConfig
	Redis *redis.Conf
}

func InitConf() *Conf {
	confOnce.Do(func() {
		conf = &Conf{
			App:   initAppConf(),
			DB:    initDBConf(),
			Redis: initRedisConf(),
		}
	})
	return conf
}
