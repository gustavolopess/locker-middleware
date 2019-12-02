package store

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/identity"
	"github.com/gustavolopess/locker/src/identity/store/postgres"
)

type IdentityStore interface {
	CreateIdentity(identity identity.Identity) error
	GetIdentityByID(id string) (identity.Identity, error)
	GetIdentityByFingerprint(fingerprint string) (identity.Identity, error)
}

type identityStore struct{
	logger log.Logger
	database postgres.IdentityDatabase
}

func NewIdentityStore(logger log.Logger) IdentityStore {
	return &identityStore{
		logger:   logger,
		database: postgres.NewIdentityDatabase(logger),
	}
}

func (i identityStore) CreateIdentity(identity identity.Identity) error {
	return i.database.CreateIdentity(identity)
}

func (i identityStore) GetIdentityByID(id string) (identity.Identity, error) {
	return i.database.GetIdentityByID(id)
}

func (i identityStore) GetIdentityByFingerprint(fingerprint string) (identity.Identity, error) {
	return i.database.GetIdentityByFingerprint(fingerprint)
}
