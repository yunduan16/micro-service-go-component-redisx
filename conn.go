package redisx

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type RedisClusterConf struct {
	Addrs []string
	RedisCommonConf
}

type RedisConf struct {
	Addr string
	RedisCommonConf
}

type RedisCommonConf struct {
	Password     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	DialTimeout  time.Duration
	MaxRetries   int
	PoolSize     int
	DB           int
}

func InitRedisCluster(conf *RedisClusterConf) (*redis.ClusterClient, func(), error) {
	options := &redis.ClusterOptions{
		Addrs:        conf.Addrs,
		Password:     conf.Password,
		PoolSize:     conf.PoolSize,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		IdleTimeout:  conf.IdleTimeout,
		MaxRetries:   conf.MaxRetries,
	}
	client := redis.NewClusterClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, nil, err
	}
	return client, func() {
		client.Close()
	}, nil
}

func InitRedis(conf *RedisConf) (*redis.Client, func(), error) {
	options := &redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		PoolSize:     conf.PoolSize,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		IdleTimeout:  conf.IdleTimeout,
		MaxRetries:   conf.MaxRetries,
		DB:           conf.DB,
	}
	client := redis.NewClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, nil, err
	}

	return client, func() {
		client.Close()
	}, nil
}
