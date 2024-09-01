package redis

type RedisDB int

const (
	RedisDBCache RedisDB = 1
	RedisDBLock  RedisDB = 2
)

type Conf struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
