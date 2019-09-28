package daemon

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port          int
	UnixSocket    string
	UseSocketFile bool
}

func InitConfig() (*Config, error) {
	viper.SetDefault("CDGO_DAEMON_PORT", 7979)
	viper.SetDefault("CDGO_DAEMON_UNIX_SOCKET", "/var/run/cdgo.sock")
	viper.SetDefault("CDGO_DAEMON_USE_UNIX_SOCKET", false)
	config := &Config{
		Port:          viper.GetInt("CDGO_DAEMON_PORT"),
		UnixSocket:    viper.GetString("CDGO_DAEMON_UNIX_SOCKET"),
		UseSocketFile: viper.GetBool("CDGO_DAEMON_USE_UNIX_SOCKET"),
	}
	return config, nil
}
