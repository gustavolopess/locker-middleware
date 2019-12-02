package store

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/locker"
	"github.com/gustavolopess/locker/src/locker/store/postgres"
	"github.com/gustavolopess/locker/src/locker/store/redis"
)

// Store describes the persistence on Locker model
type Store interface {
	CreateLocker(locker locker.Locker) (string, error)
	GetLockerByID(id string) (locker.Locker, error)
	CloseLocker(id string) error
	OpenLocker(id string) error
	IsLockerOpened(id string) bool
}

type store struct {
	logger   log.Logger
	cache    redis.LockerCache
	database postgres.LockerDatabase
}

func NewStore(logger log.Logger) Store {
	database := postgres.NewLockerDatabase(logger)
	cache := redis.NewLockerCache(logger)
	return store{logger, cache, database}
}

func (s store) CreateLocker(locker locker.Locker) (string, error) {
	return s.database.CreateLocker(locker)
}

func (s store) GetLockerByID(id string) (locker.Locker, error) {
	return s.database.GetLockerByID(id)
}

func (s store) CloseLocker(id string) error {
	return s.cache.CloseLocker(id)
}

func (s store) OpenLocker(id string) error {
	return s.cache.OpenLocker(id)
}

func (s store) IsLockerOpened(id string) bool {
	return s.cache.IsLockerOpened(id)
}
