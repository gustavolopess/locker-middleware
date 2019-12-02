package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-pg/pg"
	"github.com/gustavolopess/locker/src/mom/message"
	"github.com/gustavolopess/locker/src/mom/room"
	"os"
)

type RoomDatabase interface {
	CreateRoom(room.Room) error
	DeleteRoom(string) error
	InsertMessageOnRoom(topicID, roomID, payload string) error
	GetMessageFromRoom(oomID string, indexFrom int64) (msg message.Message, err error)
}

type roomDatabase struct{
	logger log.Logger
}

var connection *pg.DB

func (r roomDatabase) GetConnection() *pg.DB {
	if connection == nil {
		connection = pg.Connect(&pg.Options{
			Addr:                  os.Getenv("POSTGRES_ADDRESS"),
			User:                  os.Getenv("POSTGRES_USER"),
			Password:              os.Getenv("POSTGRES_PASSWORD"),
			Database:              os.Getenv("POSTGRES_DATABASE"),
			PoolSize:              40,
		})
		_ = level.Info(r.logger).Log("pgdatabase", "Init pgdatabase session")
	}

	return connection
}

func NewRoomDatabase(logger log.Logger) RoomDatabase {
	return &roomDatabase{logger}
}

func (r roomDatabase) CreateRoom(room room.Room) error {
	conn := r.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	if err := tx.Insert(&room); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (r roomDatabase) DeleteRoom(roomID string) error {
	conn := r.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	roomToDelete := &room.Room{
		ID:      roomID,
	}
	if err := tx.Delete(roomToDelete); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (r roomDatabase) InsertMessageOnRoom(topicID, roomID, payload string) error {
	msg := message.Message{
		RoomID:  roomID,
		Payload: payload,
	}

	conn := r.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	if err := tx.Insert(&msg); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (r roomDatabase) GetMessageFromRoom(roomID string, indexFrom int64) (msg message.Message, err error) {
	conn := r.GetConnection()
	err = conn.Model(&msg).
		Where("room_id = ? AND index > ?", roomID, indexFrom).
		Order("index ASC").First()
	return
}
