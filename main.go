package main

import (
	"flag"
	"fmt"

	"github.com/bitcoin-trading-system/bitcoin-redis-server/config"
	"github.com/bitcoin-trading-system/bitcoin-redis-server/router"
)

func main() {
	tomlFilePath := flag.String("toml", "toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg, err := config.NewConfig(*tomlFilePath, *envFilePath)
	if err != nil {
		panic(err)
	}

	r, err := router.NewRouter(cfg)
	if err != nil {
		panic(err)
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.BaseConfig.Port)); err != nil {
		panic(err)
	}
}
