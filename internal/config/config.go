package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env"`
	Add string `yaml:"add"`
}

func ConfigInit() *Config {
	configPath := "config.yaml"
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err.Error())
	}
	return &cfg
}
