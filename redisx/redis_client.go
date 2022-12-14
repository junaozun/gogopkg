package redisx

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisType int32

const (
	RedisType_GoReids RedisType = iota
	RedisType_Redigo
)

type Config struct {
	Server string    `json:"server"`
	Auth   string    `json:"auth"`
	Index  int       `json:"index"`
	Type   RedisType `json:"type"`
}

type IClient interface {
	Ping() error
	Close() error
	Do(args ...interface{}) (string, error)

	Set(key string, value interface{}, expireTime int64) error
	Get(key string) (string, error)
	Keys(prefix string) ([]string, error)
	Del(keys ...string) error
	IsExist(key string) (int64, error)

	MSet(kv ArrayReq) error
	MGet(keys ...string) ([]string, error)

	HGet(key string, field string) (string, error)
	HSet(key string, field string, value interface{}) error
	HMSet(key string, value ArrayReq) error
	HMGet(key string, fields ...string) ([]string, error)
	HGetAll(key string) (map[string]string, error)
	HDel(key string, fields ...string) error
	HisExist(key string, field string) (bool, error)

	ZAdd(key string, score ZAddReq) (int64, error)
	ZRange(key string, start, stop int64) ([]string, error)
	ZIncrby(key string, data map[string]int) error
	ZRem(key string, mems ...string) error
	ZRevRank(key string, mem string) (int64, error)
	ZScore(key string, mem string) (float64, error)
	ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error)

	RPush(key string, value ...interface{}) error
	LRange(key string, start, end int64) ([]string, error)
	ListAll(key string) ([]string, error)
}

func NewClient(cfg Config) (IClient, error) {
	cli := newClient(cfg)
	if err := cli.Ping(); err != nil {
		return nil, fmt.Errorf("redis ping error: %v", err)
	}
	return cli, nil
}

func newClient(cfg Config) IClient {
	switch cfg.Type {
	case RedisType_GoReids:
		return newGoRedis(cfg)
	default:
		return newGoRedis(cfg)
	}
}
