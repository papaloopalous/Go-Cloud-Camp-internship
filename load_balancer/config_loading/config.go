package configloading

import (
	"fmt"
	"load_balancer/internal/messages"

	"github.com/spf13/viper"
)

const (
	ServerAddr   = "server.address"
	BackendAddrs = "backends"
	Interval     = "interval"
)

func LoadConfig() error {
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		return fmt.Errorf(messages.ErrReadConfig, err)
	}

	return nil
}

func SetParams() (serverAddr string, backendAddrs []string, interval int) {
	serverAddr = viper.GetString(ServerAddr)
	backendAddrs = viper.GetStringSlice(BackendAddrs)
	interval = viper.GetInt(Interval)
	return serverAddr, backendAddrs, interval
}
