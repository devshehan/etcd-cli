package internal

import (
	"etcd_cli_pickme/internal/config"
	"github.com/spf13/viper"
	"log"
)

func init() {
	config.LoadAppConfig()

	viper.SetConfigName("properties")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("failed to find the etcd credential file")
		return
	}
	config.EtcdCltCfg = config.EtcdClientCfg{}

}
