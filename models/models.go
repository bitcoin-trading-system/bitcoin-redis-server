package models

import (
	"github.com/bitcoin-trading-system/bitcoin-redis-server/config"
)

type RedisRepository struct {
	RedisMap map[int]IRedis
}

func NewRedisRepository(cfg config.Config) (*RedisRepository, error) {
	redisMap := make(map[int]IRedis)
	for i := 0; i < cfg.RedisConfig.IndexCount; i++ {
		redis, err := NewRedis(cfg, i)
		if err != nil {
			return nil, err
		}
		redisMap[i] = redis
	}

	return &RedisRepository{
		RedisMap: redisMap,
	}, nil
}

func (r *RedisRepository) GetRedis(index int) IRedis {
	// TODO indexのエラーハンドリング
	return r.RedisMap[index]
}
