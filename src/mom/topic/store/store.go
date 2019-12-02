package store

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/mom/topic"
	"github.com/gustavolopess/locker/src/mom/topic/store/postgres"
)

type TopicStore interface {
	Create(topic.Topic) error
	GetTopicByID(string) (topic.Topic, error)
}

type topicStore struct {
	logger log.Logger
	database postgres.TopicDatabase
}

func NewTopicStore(logger log.Logger) TopicStore {
	return &topicStore{
		logger:   logger,
		database: postgres.NewTopicDatabase(logger),
	}
}

func (t topicStore) Create(topic topic.Topic) error {
	return t.database.Create(topic)
}

func (t topicStore) GetTopicByID(topicID string) (topic.Topic, error) {
	return t.database.GetTopicByID(topicID)
}





