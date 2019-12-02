package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-pg/pg"
	"github.com/gustavolopess/locker/src/mom/topic"
	"github.com/gustavolopess/locker/src/utils/pgdatabase"
	"os"
)

type TopicDatabase interface {
	Create(topic.Topic) error
	GetTopicByID(string) (topic.Topic, error)
}

type topicDatabase struct {
	pgdatabase.PgDatabase
	logger log.Logger
}

var connection *pg.DB

func NewTopicDatabase(logger log.Logger) TopicDatabase {
	return &topicDatabase{
		PgDatabase: pgdatabase.NewPgDatabase(logger),
		logger:     logger,
	}
}

func (t topicDatabase) GetConnection() *pg.DB {
	if connection == nil {
		connection = pg.Connect(&pg.Options{
			Addr:                  os.Getenv("POSTGRES_ADDRESS"),
			User:                  os.Getenv("POSTGRES_USER"),
			Password:              os.Getenv("POSTGRES_PASSWORD"),
			Database:              os.Getenv("POSTGRES_DATABASE"),
			PoolSize:              40,
		})
		_ = level.Info(t.logger).Log("pgdatabase", "Init pgdatabase session")
	}

	return connection
}

func (t topicDatabase) Create(tpc topic.Topic) error {
	conn := t.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	if err := tx.Insert(&tpc); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (t topicDatabase) GetTopicByID(topicID string) (tpc topic.Topic, err error) {
	conn := t.GetConnection()
	err = conn.Model(&tpc).Where("id = ?", topicID).First()
	return
}
