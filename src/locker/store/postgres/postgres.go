package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/locker"
	"github.com/gustavolopess/locker/src/utils/pgdatabase"
)

type LockerDatabase interface {
	CreateLocker(locker locker.Locker) error
	GetLockerByID(id string) (locker.Locker, error)
}

type lockerDatabase struct {
	pgdatabase.PgDatabase
	logger log.Logger
}

func NewLockerDatabase(logger log.Logger) LockerDatabase {
	return &lockerDatabase{
		PgDatabase: pgdatabase.NewPgDatabase(logger),
		logger:     logger,
	}
}

func (db *lockerDatabase) CreateLocker(locker locker.Locker) error {
	return db.CreateSomething(locker)
}

func (db *lockerDatabase) GetLockerByID(id string) (locker.Locker, error) {
	loc := locker.Locker{}

	err := db.GetModelByAttribute(loc, "id", id)

	return loc, err
}



