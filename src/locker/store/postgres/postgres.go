package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-pg/pg"
	"github.com/gustavolopess/locker/src/locker"
	"os"
)

type LockerDatabase interface {
	CreateLocker(id string) (string, error)
	GetLockerByID(id string) (locker.Locker, error)
}

type lockerDatabase struct {
	logger log.Logger
}

var connection *pg.DB

func NewLockerDatabase(logger log.Logger) LockerDatabase {
	return &lockerDatabase{
		logger:     logger,
	}
}

func (db *lockerDatabase) GetConnection() *pg.DB {
	if connection == nil {
		connection = pg.Connect(&pg.Options{
			Addr:                  os.Getenv("POSTGRES_LOCKER_ADDRESS"),
			User:                  os.Getenv("POSTGRES_LOCKER_USER"),
			Password:              os.Getenv("POSTGRES_LOCKER_PASSWORD"),
			Database:              os.Getenv("POSTGRES_LOCKER_DATABASE"),
			PoolSize:              40,
		})
		_ = level.Info(db.logger).Log("locker-database", "Init pgdatabase session")
	}

	return connection
}

func (db *lockerDatabase) CreateLocker(id string) (string, error) {
	loc := locker.Locker{ID:id}

	conn := db.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return "", err
	}
	if err := tx.Insert(&loc); err != nil {
		_ = tx.Rollback()
		return "", err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return "", err
	}

	return loc.ID, nil
}

func (db *lockerDatabase) GetLockerByID(id string) (locker.Locker, error) {
	conn := db.GetConnection()

	var loc locker.Locker
	err := conn.Model(&loc).Where("id = ?", id).First()

	return loc, err
}
