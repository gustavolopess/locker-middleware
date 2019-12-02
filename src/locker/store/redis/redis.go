package redis

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"time"
)

type RedisCache interface {
	GetConnection() *redis.Client
	CloseConnection()
	GetByKey(key string) (string, error)
}

type LockerCache interface {
	OpenLocker(id string) error
	CloseLocker(id string) error
	IsLockerOpened(id string) bool
}

type redisCache struct {
	logger log.Logger
}

type lockerCache struct {
	RedisCache
	logger log.Logger
}

var connection *redis.Client

func (r redisCache) GetConnection() *redis.Client {
	if connection == nil {
		redisDatabase, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
		if err != nil {
			panic("Invalid redis database")
		}
		connection = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDatabase,
		})
		_ = level.Info(r.logger).Log("redis", "init redis connection")
	}

	return connection
}

func (r redisCache) CloseConnection() {
	if connection != nil {
		_ = level.Info(r.logger).Log("redis", "closing redis connection")
		_ = connection.Close()
	}
}

func NewLockerCache(logger log.Logger) LockerCache {
	return lockerCache{
		RedisCache: redisCache{logger},
		logger:     logger,
	}
}

func (r redisCache) GetByKey(key string) (string, error) {
	conn := r.GetConnection()
	return conn.Get(key).Result()
}

func (l lockerCache) OpenLocker(id string) error {
	conn := l.GetConnection()
	return conn.Set(id, "open", 20 * time.Second).Err()
}

func (l lockerCache) CloseLocker(id string) error {
	conn := l.GetConnection()
	return conn.Del(id).Err()
}

func (l lockerCache) IsLockerOpened(id string) bool {
	status, err := l.GetByKey(id)
	if err != nil {
		return false
	}

	return status == "open"
}