package app

import "github.com/spf13/viper"

type Config struct {
	Host string
	Port int
	Nick string
}

func InitConfig() (*Config, error) {
	viper.SetDefault("CDGO_NODE_HOST", "127.0.0.1")
	viper.SetDefault("CDGO_NODE_PORT", "7777")
	config := &Config{
		Host: viper.GetString("CDGO_NODE_HOST"),
		Port: viper.GetInt("CDGO_NODE_PORT"),
		Nick: viper.GetString("CDGO_USER_NICK"),
	}
	return config, nil
}
