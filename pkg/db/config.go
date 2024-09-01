package db

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Keyspace string `mapstructure:"keyspace"`
}
