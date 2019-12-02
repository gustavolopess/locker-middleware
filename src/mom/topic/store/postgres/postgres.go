package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/mom/topic"
	"github.com/gustavolopess/locker/src/utils/pgdatabase"
)

type TopicDatabase interface {
	Create(topic.Topic) error
	GetTopicByID(string) (topic.Topic, error)
}

type topicDatabase struct {
	pgdatabase.PgDatabase
	logger log.Logger
}

func NewTopicDatabase(logger log.Logger) TopicDatabase {
	return &topicDatabase{
		PgDatabase: pgdatabase.NewPgDatabase(logger),
		logger:     logger,
	}
}

//func (t topicDatabase) isRoomValid(topicID, roomID string) bool {
//	var r room.Room
//	if err := t.GetModelByAttribute(&r, "room_id", roomID); err != nil {
//		return false
//	}
//
//	return r.TopicID != topicID
//}

func (t topicDatabase) Create(tpc topic.Topic) error {
	return t.CreateSomething(tpc)
}

func (t topicDatabase) GetTopicByID(topicID string) (tpc topic.Topic, err error) {
	err = t.GetModelByAttribute(&tpc, "topic_id", topicID)
	return
}
