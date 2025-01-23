package models

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/bitcoin-trading-system/bitcoin-redis-server/config"
)

var ctx = context.Background()

type Redis struct {
	Client *redis.Client
	Config config.Config
}

type IRedis interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
}

func NewRedis(cfg config.Config, index int) (IRedis, error) {
	if index < 0 || index >= cfg.RedisConfig.IndexCount {
		return nil, errors.New("index is out of range")
	}

	client, err := ConnectRedis(cfg, index)
	if err != nil {
		return nil, err
	}

	return &Redis{
		Client: client,
		Config: cfg,
	}, nil
}

func ConnectRedis(cfg config.Config, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Address,
		Password: "",
		DB:       db,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}

func (r *Redis) Set(key string, value interface{}, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *Redis) Del(key string) error {
	return r.Client.Del(ctx, key).Err()
}
