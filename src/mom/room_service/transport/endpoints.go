package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gustavolopess/locker/src/mom/room"
	"github.com/gustavolopess/locker/src/mom/room_service"
)

// RoomEndpoints holds all go kit's endpoints for room service
type RoomEndpoints struct {
	CreateRoom          endpoint.Endpoint
	DeleteRoom          endpoint.Endpoint
	InsertMessageOnRoom endpoint.Endpoint
	GetMessageFromRoom  endpoint.Endpoint
}

func MakeRoomEndpoints(svc room_service.RoomService) RoomEndpoints {
	return RoomEndpoints{
		CreateRoom:          makeCreateRoomEndpoint(svc),
		DeleteRoom:          makeDeleteRoomEndpoint(svc),
		InsertMessageOnRoom: makeInsertMessageOnRoomEndpoint(svc),
		GetMessageFromRoom:  makeGetMessageFromRoomEndpoint(svc),
	}
}

func makeCreateRoomEndpoint(svc room_service.RoomService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRoomRequest)
		newRoom := room.Room{
			TopicID: req.TopicID,
		}
		err := svc.CreateRoom(ctx, newRoom)
		return CreateRoomResponse{err}, nil
	}
}

func makeDeleteRoomEndpoint(svc room_service.RoomService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRoomRequest)
		err := svc.DeleteRoom(ctx, req.RoomID)
		return DeleteRoomResponse{err}, nil
	}
}

func makeInsertMessageOnRoomEndpoint(svc room_service.RoomService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertMessageOnRoomRequest)
		err := svc.InsertMessageOnRoom(ctx, req.TopicID, req.RoomID, req.Payload)
		return InsertMessageOnRoomResponse{err}, nil
	}
}

func makeGetMessageFromRoomEndpoint(svc room_service.RoomService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMessageFromRoomRequest)
		msg, err := svc.GetMessageFromRoom(ctx, req.RoomID, req.Index)
		return GetMessageFromRoomResponse{
			Message: msg,
			Err:     err,
		}, nil
	}
}
