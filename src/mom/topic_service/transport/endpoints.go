package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gustavolopess/locker/src/mom/topic"
	"github.com/gustavolopess/locker/src/mom/topic_service"
)

// TopicEndpoints holds Go endpoints for TopicService
type TopicEndpoints struct {
	CreateTopic endpoint.Endpoint
	GetTopicByID endpoint.Endpoint
}

func MakeTopicEndpoints(svc topic_service.TopicService) TopicEndpoints {
	return TopicEndpoints{
		CreateTopic:  makeCreateTopicEndpoint(svc),
		GetTopicByID: makeGetTopicByIDEndpoint(svc),
	}
}

func makeCreateTopicEndpoint(svc topic_service.TopicService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTopicRequest)
		tpc := topic.Topic{
			Name: req.Name,
		}
		err := svc.CreateTopic(ctx, tpc)
		return CreateTopicResponse{err}, nil
	}
}

func makeGetTopicByIDEndpoint(svc topic_service.TopicService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTopicByIDRequest)
		tpc, err := svc.GetTopicByID(ctx, req.ID)
		return GetTopicByIDResponse{tpc, err}, nil
	}
}