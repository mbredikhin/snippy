package repository

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisDB(config RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	return redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: config.Password,
			DB:       config.DB,
		},
	)
}
