package config

import "github.com/jinzhu/configor"

type Config struct {
	Port string `default:"1323"`
}

func New() *Config {
	config := new(Config)

	err := configor.Load(config, "./config/config.yaml")
	if err != nil {
		return nil
	}

	return config
}
