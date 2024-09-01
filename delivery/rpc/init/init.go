package init

import (
	"sync"
)

var conf *Conf
var confOnce sync.Once

type Conf struct {
	App *App
}

func InitConf() *Conf {
	confOnce.Do(func() {
		conf = &Conf{
			App: initAppConf(),
		}
	})
	return conf
}
