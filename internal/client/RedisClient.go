package client

import (
	"context"
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/config"
	"github.com/redis/go-redis/v9"
	"log"
)

func RedisClientConnect(cfg config.Redis) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.Db,
	}
	redisClient := redis.NewClient(opts)
	if status := redisClient.Ping(context.Background()); status.Err() != nil {
		return nil, status.Err()
	}
	log.Println("redis connect success")
	return redisClient, nil
}

func NewClientRedis() *redis.Client {
	client, err := RedisClientConnect(global.EngineConfig.Redis)

	if err != nil {
		log.Println("redis connect fail")
	}
	return client
}
