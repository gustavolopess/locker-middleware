package network

import (
	"encoding/json"
)

const (
	CreateTopicOperation         = "CREATE_TOPIC"
	GetTopicByIDOperation        = "GET_TOPIC_BY_ID"
	CreateRoomOperation          = "CREATE_ROOM"
	DeleteRoomOperation          = "DELETE_ROOM"
	InsertMessageOnRoomOperation = "INSERT_MESSAGE_ON_ROOM"
	GetMessageFromRoomOperation  = "GET_MESSAGE_FROM_ROOM"
)

type GenericRequest struct {
	Operation string `json:"operation"`
	TopicID   string `json:"topic_id,omitempty"`
	TopicName string `json:"topic_name,omitempty"`
	RoomID    string `json:"room_id,omitempty"`
	Payload   string `json:"payload,omitempty"`
	Index     int64  `json:"index,omitempty"`
}

func ByteToRequest(data []byte) (req GenericRequest, err error) {
	if err = json.Unmarshal(data, &req); err != nil {
		return
	}
	return
}

func RequestToByte(req GenericRequest) ([]byte, error) {
	return json.Marshal(req)
}

//func MarshalRequest(req []byte) (req GenericRequest, err error) {
//	req.Payload = v
//
//	switch v.(type) {
//	case transport.CreateRoomRequest:
//		req.Operation = CreateRoomOperation
//		break
//	case transport.DeleteRoomRequest:
//		req.Operation = DeleteRoomOperation
//		break
//	case transport.InsertMessageOnRoomRequest:
//		req.Operation = InsertMessageOnRoomOperation
//		break
//	case transport.GetMessageFromRoomRequest:
//		req.Operation = GetMessageFromRoomOperation
//		break
//	default:
//		err = errors.New("invalid request")
//	}
//
//	return
//}
//
//func UnmarshalRequest(req GenericRequest) (v interface{}, err error) {
//	switch req.Operation {
//	case CreateRoomOperation:
//		v = v.(transport.CreateRoomRequest)
//
//	}
//}
