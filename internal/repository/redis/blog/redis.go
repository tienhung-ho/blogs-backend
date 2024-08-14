package blogcachestorage

import "github.com/redis/go-redis/v9"

type redisStorage struct {
	rdb *redis.Client
}

func NewRedisStorage(rdb *redis.Client) *redisStorage {
	return &redisStorage{rdb: rdb}
}
