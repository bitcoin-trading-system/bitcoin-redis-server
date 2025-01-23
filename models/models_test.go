package models

import "github.com/bitcoin-trading-system/bitcoin-redis-server/config"

var modelsTestConfig config.Config

const redisTestIndex = 0

func init() {
	var err error
	modelsTestConfig, err = config.NewConfig("../toml/local.toml", "")
	if err != nil {
		panic(err)
	}
}
