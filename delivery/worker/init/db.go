package init

import (
	"github.com/tony-zhuo/follow-service/pkg/config"
	"github.com/tony-zhuo/follow-service/pkg/db"
)

func initDBConf() *db.DBConfig {
	path := "./conf"
	name := "db"
	configType := "yml"
	conf := &db.DBConfig{}
	config.Load(path, name, configType, conf)
	return conf
}
