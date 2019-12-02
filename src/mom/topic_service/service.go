package topic_service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/mom/topic"
	"github.com/gustavolopess/locker/src/mom/topic/store"
)

type TopicService interface {
	CreateTopic(context.Context, topic.Topic) error
	GetTopicByID(context.Context, string) (topic.Topic, error)
}

type topicService struct {
	store store.TopicStore
	logger log.Logger
}

func NewTopicService(store store.TopicStore, logger log.Logger) TopicService {
	return &topicService{
		store: store,
		logger: logger,
	}
}

func (t topicService) CreateTopic(ctx context.Context, tpc topic.Topic) error {
	return t.store.Create(tpc)
}

func (t topicService) GetTopicByID(ctx context.Context, topicID string) (topic.Topic, error) {
	return t.store.GetTopicByID(topicID)
}

