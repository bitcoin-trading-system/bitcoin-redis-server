package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

type Config struct {
	BaseConfig  BaseConfig  `toml:"baseConfig"`
	RedisConfig RedisConfig `toml:"redisConfig"`
}

type BaseConfig struct {
	Port string `toml:"port"`
}

type RedisConfig struct {
	Address    string `toml:"address"`
	IndexCount int    `toml:"indexCount"`
}

func NewConfig(tomlFilePath, envFilePath string) (Config, error) {
	var cfg *Config

	if _, err := toml.DecodeFile(tomlFilePath, &cfg); err != nil {
		panic(err)
	}

	return *cfg, cfg.mustCheck()
}

func (cfg *Config) mustCheck() error {
	if cfg.BaseConfig.Port == "" {
		return errors.New("host is required")
	}

	if cfg.RedisConfig.Address == "" {
		return errors.New("address is required")
	}

	if cfg.RedisConfig.IndexCount == 0 {
		return errors.New("indexCount is required")
	}

	return nil
}
