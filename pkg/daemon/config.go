package daemon

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port int
	Nick string
}

func InitConfig() (*Config, error) {
	config := &Config{
		Host: viper.GetString("CDGO_DAEMON_HOST"),
		Port: viper.GetInt("CDGO_DAEMON_PORT"),
		Nick: viper.GetString("CDGO_USER_NICK"),
	}
	return config, nil
}
