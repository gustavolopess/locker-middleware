package locker_service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/identity"
	"github.com/gustavolopess/locker/src/locker"
	lockerstore "github.com/gustavolopess/locker/src/locker/store"
)

type Service interface {
	CreateLocker(ctx context.Context, locker locker.Locker) (string, error)
	GetLockerByID(ctx context.Context, id string) (locker.Locker, error)
	OpenLocker(ctx context.Context, id string, fingerprint identity.Fingerprint) error
	CloseLocker(ctx context.Context, id string) error
}


type service struct {
	store  lockerstore.Store
	logger log.Logger
}

func NewService(store lockerstore.Store, logger log.Logger) Service {
	return &service{
		store:  store,
		logger: logger,
	}
}

func (s *service) CreateLocker(ctx context.Context, locker locker.Locker) (string, error) {
	return s.store.CreateLocker(locker)
}

func (s *service) GetLockerByID(ctx context.Context, id string) (locker.Locker, error) {
	return s.store.GetLockerByID(id)
}

func (s *service) OpenLocker(ctx context.Context, id string, fingerprint identity.Fingerprint) error {
	// TODO: validate fingerprint before open locker
	return s.store.OpenLocker(id)
}

func (s *service) CloseLocker(ctx context.Context, id string) error {
	return s.store.CloseLocker(id)
}
