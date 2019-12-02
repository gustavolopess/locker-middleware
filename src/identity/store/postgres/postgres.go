package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/identity"
	"github.com/gustavolopess/locker/src/utils/pgdatabase"
)

type IdentityDatabase interface {
	CreateIdentity(identity identity.Identity) error
	GetIdentityByID(id string) (identity.Identity, error)
	GetIdentityByFingerprint(fingerprint identity.Fingerprint) (identity.Identity, error)
}

type identityDatabase struct{
	pgdatabase.PgDatabase
	logger log.Logger
}

func NewIdentityDatabase(logger log.Logger) IdentityDatabase {
	return &identityDatabase{
		PgDatabase: pgdatabase.NewPgDatabase(logger),
		logger:     logger,
	}
}

func (db *identityDatabase) CreateIdentity(identity identity.Identity) error {
	return db.CreateSomething(identity)
}

func (db *identityDatabase) GetIdentityByID(id string) (identity.Identity, error) {
	ident := identity.Identity{}

	err := db.GetModelByAttribute(ident, "id", id)

	return ident, err
}

func (db *identityDatabase) GetIdentityByFingerprint(fingerprint identity.Fingerprint) (identity.Identity, error) {
	ident := identity.Identity{}

	err := db.GetModelByAttribute(ident, "fingerprint", fingerprint)

	return ident, err
}

