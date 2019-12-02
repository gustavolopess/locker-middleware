package main

import (
	"github.com/gustavolopess/locker/src/mom/network"
	roomStore "github.com/gustavolopess/locker/src/mom/room/store"
	"github.com/gustavolopess/locker/src/mom/room_service"
	roomTransport "github.com/gustavolopess/locker/src/mom/room_service/transport"
	topicStore "github.com/gustavolopess/locker/src/mom/topic/store"
	"github.com/gustavolopess/locker/src/mom/topic_service"
	topicTransport "github.com/gustavolopess/locker/src/mom/topic_service/transport"
	logger2 "github.com/gustavolopess/locker/src/utils/logger"
)

func buildRoomEndpoints() roomTransport.RoomEndpoints {
	logger := logger2.BuildLogger("mom-room")

	store := roomStore.NewRoomStore(logger)
	roomSvc := room_service.NewRoomService(store, logger)

	return roomTransport.MakeRoomEndpoints(roomSvc)
}

func buildTopicEndpoints() topicTransport.TopicEndpoints {
	logger := logger2.BuildLogger("mom-topic")

	store := topicStore.NewTopicStore(logger)
	topicService := topic_service.NewTopicService(store, logger)

	return topicTransport.MakeTopicEndpoints(topicService)
}

func main() {
	topicEndpoints := buildTopicEndpoints()
	roomEndpoints := buildRoomEndpoints()

	logger := logger2.BuildLogger("server")
	server := network.NewServer(8089, logger)

	if err := server.Listen(topicEndpoints, roomEndpoints); err != nil {
		panic(err)
	}
}
