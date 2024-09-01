package config

import (
	"github.com/spf13/viper"
)

func Load(configDir, filename, configType string, config interface{}) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(configType)
	v.AddConfigPath(configDir)
	err := v.ReadInConfig()
	if err != nil {
		// logging
	}

	v.Unmarshal(&config)
}
