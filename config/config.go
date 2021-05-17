package config

import "github.com/jinzhu/configor"

type Config struct {
	Port  string `default:"1323"`
	Redis struct {
		Password string `default:""`
		Port     string `default:"6379"`
		Host     string `default:"127.0.0.1"`
	}
}

func New() *Config {
	config := new(Config)

	err := configor.Load(config, "./config/config.yaml")
	if err != nil {
		return nil
	}

	return config
}
