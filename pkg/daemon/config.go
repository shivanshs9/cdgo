package daemon

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port int
}

func InitConfig() (*Config, error) {
	config := &Config{
		Host: viper.GetString("CDGO_DAEMON_HOST"),
		Port: viper.GetInt("CDGO_DAEMON_PORT"),
	}
	return config, nil
}
