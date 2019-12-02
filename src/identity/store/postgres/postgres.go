package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-pg/pg"
	"github.com/gustavolopess/locker/src/identity"
	"os"
)

type IdentityDatabase interface {
	CreateIdentity(identity identity.Identity) error
	GetIdentityByID(id string) (identity.Identity, error)
	GetIdentityByFingerprint(fingerprint string) (identity.Identity, error)
}

type identityDatabase struct{
	logger log.Logger
}

var connection *pg.DB

func NewIdentityDatabase(logger log.Logger) IdentityDatabase {
	return &identityDatabase{
		logger:     logger,
	}
}

func (db *identityDatabase) GetConnection() *pg.DB {
	if connection == nil {
		connection = pg.Connect(&pg.Options{
			Addr:                  os.Getenv("POSTGRES_LOCKER_ADDRESS"),
			User:                  os.Getenv("POSTGRES_LOCKER_USER"),
			Password:              os.Getenv("POSTGRES_LOCKER_PASSWORD"),
			Database:              os.Getenv("POSTGRES_LOCKER_DATABASE"),
			PoolSize:              40,
		})
		_ = level.Info(db.logger).Log("identity-database", "Init pgdatabase session")
	}

	return connection
}

func (db *identityDatabase) CreateIdentity(identity identity.Identity) error {
	conn := db.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	if err := tx.Insert(&identity); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (db *identityDatabase) GetIdentityByID(id string) (idt identity.Identity, err error) {
	conn := db.GetConnection()
	err = conn.Model(&idt).Where("id = ?", id).First()
	return
}

func (db *identityDatabase) GetIdentityByFingerprint(fingerprint string) (idt identity.Identity, err error) {
	conn := db.GetConnection()
	err = conn.Model(&idt).Where("fingerprint = ?", fingerprint).First()
	return
}

