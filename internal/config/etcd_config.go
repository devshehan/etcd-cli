package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var EtcdCltCfg EtcdClientCfg

type EtcdClientCfg struct {
	Server   string `yaml:"server"`
	Port     int64  `yaml:"port"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

func LoadETCDConfig(env string) (err error) {
	root := viper.Sub("etcd-root")
	if root == nil {
		return errors.New("failed to find 'etcd-root'")
	}

	envConfig := root.Sub(env)
	if envConfig == nil {
		return errors.New(fmt.Sprintf("failed to find the section %s", env))
	}

	if err = envConfig.Unmarshal(&EtcdCltCfg); err != nil {
		return errors.New(fmt.Sprintf("failed to unmarshell envconfig"))
	}
	return nil
}
