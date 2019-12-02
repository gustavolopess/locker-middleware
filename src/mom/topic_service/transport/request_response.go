package transport

import (
	"github.com/gustavolopess/locker/src/mom/topic"
)

type CreateTopicRequest struct {
	Name string
}

type CreateTopicResponse struct {
	Err error
}

type GetTopicByIDRequest struct {
	ID string
}

type GetTopicByIDResponse struct {
	Topic topic.Topic
	Err error
}
