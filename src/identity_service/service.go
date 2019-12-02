package identity_service

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/identity"
	"github.com/gustavolopess/locker/src/identity/store"
)

type IdentityService interface {
	CreateIdentity(identity identity.Identity) error
	GetIdentityByID(id string) (identity.Identity, error)
	GetIdentityByFingerprint(fingerprint string) (identity.Identity, error)
}

type identityService struct{
	store store.IdentityStore
	logger log.Logger
}

func NewIdentityService(logger log.Logger) IdentityService {
	return &identityService{
		store: store.NewIdentityStore(logger),
		logger: logger,
	}
}

func (i identityService) CreateIdentity(identity identity.Identity) error {
	return i.store.CreateIdentity(identity)
}

func (i identityService) GetIdentityByID(id string) (identity.Identity, error) {
	return i.store.GetIdentityByID(id)
}

func (i identityService) GetIdentityByFingerprint(fingerprint string) (identity.Identity, error) {
	return i.store.GetIdentityByFingerprint(fingerprint)
}
