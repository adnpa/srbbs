package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"srbbs/src/conf"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd

//type StringStringMapCmd = redis.StringStringMapCmd

func init() {
	cfg := conf.Cfg.RedisConfig
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:           cfg.DB,
		Password:     cfg.Password,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	panic(client.Ping(context.TODO()).Err())
}

func Close() {
	_ = client.Close()
}
