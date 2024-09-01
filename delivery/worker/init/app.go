package init

import "github.com/tony-zhuo/follow-service/pkg/config"

type App struct {
	FollowConsumer *FollowConsumerConf `mapstructure:"follow_consumer"`
}

type FollowConsumerConf struct {
	Name           string   `mapstructure:"name"`
	Enable         bool     `mapstructure:"enable"`
	KafkaServerUrl []string `mapstructure:"kafka_server_url"`
	SASLEnable     bool     `mapstructure:"sasl_enable"`
	UserName       string   `mapstructure:"user_name"`
	Password       string   `mapstructure:"password"`
	WorkerCount    int      `mapstructure:"worker_count"`
	Topic          string   `mapstructure:"topic"`
	Group          string   `mapstructure:"group"`
}

func initAppConf() *App {
	path := "./conf"
	name := "app"
	configType := "yml"
	conf := &App{}
	config.Load(path, name, configType, conf)
	return conf
}
