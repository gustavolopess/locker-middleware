package room_service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/mom/message"
	"github.com/gustavolopess/locker/src/mom/room"
	"github.com/gustavolopess/locker/src/mom/room/store"
)

type RoomService interface {
	CreateRoom(context.Context, room.Room) error
	DeleteRoom(context.Context, string) error
	InsertMessageOnRoom(ctx context.Context, topicID, roomID, payload string) error
	GetMessageFromRoom(ctx context.Context, roomID string, indexFrom int64) (msg message.Message, err error)
}

type roomService struct{
	logger log.Logger
	store store.RoomStore
}

func NewRoomService(store store.RoomStore, logger log.Logger) RoomService {
	return &roomService{
		logger: logger,
		store:  store,
	}
}

func (r roomService) CreateRoom(ctx context.Context, room room.Room) error {
	return r.store.CreateRoom(room)
}

func (r roomService) DeleteRoom(ctx context.Context, roomID string) error {
	return r.store.DeleteRoom(roomID)
}

func (r roomService) InsertMessageOnRoom(ctx context.Context, topicID, roomID, payload string) error {
	return r.store.InsertMessageOnRoom(topicID, roomID, payload)
}

func (r roomService) GetMessageFromRoom(ctx context.Context, roomID string, indexFrom int64) (msg message.Message, err error) {
	return r.store.GetMessageFromRoom(roomID, indexFrom)
}
