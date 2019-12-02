package transport

import "github.com/gustavolopess/locker/src/mom/message"

type CreateRoomRequest struct {
	TopicID string `json:"topic_id"`
}

type CreateRoomResponse struct {
	Err error `json:"error"`
}

type DeleteRoomRequest struct {
	RoomID string `json:"room_id"`
}

type DeleteRoomResponse struct {
	Err error `json:"error"`
}

type InsertMessageOnRoomRequest struct {
	TopicID string `json:"topic_id"`
	RoomID string `json:"room_id"`
	Payload string `json:"payload"`
}

type InsertMessageOnRoomResponse struct {
	Err error `json:"error"`
}

type GetMessageFromRoomRequest struct {
	RoomID string `json:"room_id"`
	Index int64    `json:"index"`
}

type GetMessageFromRoomResponse struct {
	Message message.Message `json:"message"`
	Err interface{} `json:"error"`
}