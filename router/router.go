package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-system/bitcoin-redis-server/config"
	"github.com/bitcoin-trading-system/bitcoin-redis-server/handler"
)

func NewRouter(cfg config.Config) (*gin.Engine, error) {
	h, err := handler.NewHandler(cfg)
	if err != nil {
		return nil, err
	}

	r := gin.Default()

	r.GET("index/:index/key/:key", h.GetRedisHandler)
	r.POST("index/:index/key/:key", h.PostRedisHandler)
	r.DELETE("index/:index/key/:key", h.DeleteRedisHandler)

	return r, nil
}
