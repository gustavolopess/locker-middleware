package pgdatabase

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-pg/pg"
	"os"
)

var (
	ErrOnBeginTransaction = errors.New("begin of transaction failed")
	ErrOnCommit = errors.New("transaction commit failed")
	ErrOnSelect = errors.New("select failed")
	ErrNotFound = errors.New("not found")
	ErrOnDelete = errors.New("deletion failed")
)

type PgDatabase interface {
	GetConnection() *pg.DB
	CloseConnection()
	CreateSomething(interface{}) error
	DeleteSomething(interface{}) error
	GetModelByAttribute(interface{}, string, ...interface{}) error
	GetModelByCustomQuery(interface{}, string, ...interface{}) error
}

type pgDatabase struct {
	logger log.Logger
}

var connection *pg.DB

func NewPgDatabase(logger log.Logger) PgDatabase {
	return &pgDatabase{logger: logger}
}

func (db pgDatabase) GetConnection() *pg.DB {
	if connection == nil {
		connection = pg.Connect(&pg.Options{
			Addr:                  os.Getenv("POSTGRES_ADDRESS"),
			User:                  os.Getenv("POSTGRES_USER"),
			Password:              os.Getenv("POSTGRES_PASSWORD"),
			Database:              os.Getenv("POSTGRES_DATABASE"),
			PoolSize:              40,
		})
		_ = level.Info(db.logger).Log("pgdatabase", "Init pgdatabase session")
	}

	return connection
}

func (db pgDatabase) CloseConnection() {
	if connection != nil {
		_ = level.Info(db.logger).Log("pgdatabase", "Closing pgdatabase sesison")
		_ = connection.Close()
	}
}

func (db pgDatabase) CreateSomething(something interface{}) error {
	conn := db.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	if err := tx.Insert(something); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return ErrOnCommit
	}

	return nil
}

func (db pgDatabase) DeleteSomething(something interface{}) error {
	conn := db.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return ErrOnBeginTransaction
	}

	if err := tx.Delete(something); err != nil {
		_ = tx.Rollback()
		return ErrOnDelete
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return ErrOnCommit
	}

	return nil
}

func (db pgDatabase) GetModelByAttribute(model interface{}, attribute string, params ...interface{}) error {
	conn := db.GetConnection()

	err := conn.Model(&model).Where(fmt.Sprintf("%s = ?", attribute), params).Select()

	if err == pg.ErrNoRows {
		return ErrNotFound
	}

	if err != nil {
		return ErrOnSelect
	}

	return nil
}

func (db pgDatabase) GetModelByCustomQuery(model interface{}, query string, params ...interface{}) error {
	conn := db.GetConnection()

	err := conn.Model(&model).Where(query, params).Select()

	if err == pg.ErrNoRows {
		return ErrNotFound
	}

	if err != nil {
		return ErrOnSelect
	}

	return nil
}