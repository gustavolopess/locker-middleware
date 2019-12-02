package locker_service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/locker"
	lockerstore "github.com/gustavolopess/locker/src/locker/store"
)

type LockerService interface {
	CreateLocker(ctx context.Context, id string) (string, error)
	GetLockerByID(ctx context.Context, id string) (locker.Locker, error)
	OpenLocker(ctx context.Context, id string) error
	CloseLocker(ctx context.Context, id string) error
}


type service struct {
	store  lockerstore.LockerStore
	logger log.Logger
}

func NewLockerService(store lockerstore.LockerStore, logger log.Logger) LockerService {
	return &service{
		store:  store,
		logger: logger,
	}
}

func (s *service) CreateLocker(ctx context.Context, id string) (string, error) {
	return s.store.CreateLocker(id)
}

func (s *service) GetLockerByID(ctx context.Context, id string) (locker.Locker, error) {
	return s.store.GetLockerByID(id)
}

func (s *service) OpenLocker(ctx context.Context, id string) error {
	return s.store.OpenLocker(id)
}

func (s *service) CloseLocker(ctx context.Context, id string) error {
	return s.store.CloseLocker(id)
}
