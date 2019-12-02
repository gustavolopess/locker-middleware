package store

import (
	"github.com/go-kit/kit/log"
	"github.com/gustavolopess/locker/src/mom/message"
	"github.com/gustavolopess/locker/src/mom/room"
	"github.com/gustavolopess/locker/src/mom/room/store/postgres"
)

type RoomStore interface {
	CreateRoom(room.Room) error
	DeleteRoom(string) error
	InsertMessageOnRoom(topicID, roomID, payload string) error
	GetMessageFromRoom(roomID string, indexFrom int64) (msg message.Message, err error)
}

type roomStore struct{
	logger log.Logger
	database postgres.RoomDatabase
}

func NewRoomStore(logger log.Logger) RoomStore {
	return &roomStore{
		logger:   logger,
		database: postgres.NewRoomDatabase(logger),
	}
}

func (r roomStore) CreateRoom(room room.Room) error {
	return r.database.CreateRoom(room)
}

func (r roomStore) DeleteRoom(roomID string) error {
	return r.database.DeleteRoom(roomID)
}

func (r roomStore) InsertMessageOnRoom(topicID, roomID, payload string) error {
	return r.database.InsertMessageOnRoom(topicID, roomID, payload)
}

func (r roomStore) GetMessageFromRoom(roomID string, indexFrom int64) (msg message.Message, err error) {
	return r.database.GetMessageFromRoom(roomID, indexFrom)
}


