package network

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	transport2 "github.com/gustavolopess/locker/src/mom/room_service/transport"
	"github.com/gustavolopess/locker/src/mom/topic_service/transport"
	"net"
)

type Server interface {
	Listen(topicEndpoints transport.TopicEndpoints, roomEndpoints transport2.RoomEndpoints) error
	Send(conn net.Conn, payload []byte) error
}

type server struct {
	Port   int
	Logger log.Logger
}

func NewServer(port int, logger log.Logger) Server {
	return &server{
		Port:   port,
		Logger: logger,
	}
}

func (s server) Listen(topicEndpoints transport.TopicEndpoints, roomEndpoints transport2.RoomEndpoints) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	defer listener.Close()
	if err != nil {
		return err
	}

	_ = level.Info(s.Logger).Log("server-listen", fmt.Sprintf("listening on port %d", s.Port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			_ = level.Debug(s.Logger).Log("server-listen", err.Error())
			continue
		}

		remoteHost := conn.RemoteAddr().(*net.TCPAddr).IP.String()
		_ = level.Info(s.Logger).
			Log("server-listen", fmt.Sprintf("received connection from %s", remoteHost))

		go func() {
			defer conn.Close()
			for {
				msg, err := bufio.NewReader(conn).ReadBytes('\n')
				if err != nil {
					_ = level.
						Info(s.Logger).
						Log("server-listen", fmt.Sprintf("%s disconnected", remoteHost))
					_ = level.Debug(s.Logger).Log("server-listen", err.Error())
					return
				}

				req, err := ByteToRequest(msg)

				if err != nil {
					_ = level.Debug(s.Logger).Log("server-listen", err)
					return
				}

				response, err := invokeEndpoint(req, topicEndpoints, roomEndpoints)
				if err != nil {
					_ = level.Debug(s.Logger).Log("server-listen", err)
					return
				}

				bytesResp, err := json.Marshal(response)
				if err != nil {
					_ = level.Debug(s.Logger).Log("server-listen", err)
					return
				}

				bytesResp = append(bytesResp, '\n')
				if err = s.Send(conn, bytesResp); err != nil {
					_ = level.Debug(s.Logger).Log("server-listen", err)
					return
				}

				_ = level.Info(s.Logger).Log("server-listen", fmt.Sprintf("%s processed", req.Operation))
			}
		}()
	}

}

func invokeEndpoint(req GenericRequest,
	topicEndpoints transport.TopicEndpoints, roomEndpoints transport2.RoomEndpoints) (resp interface{}, err error) {
	switch req.Operation {
	case CreateTopicOperation:
		return topicEndpoints.CreateTopic(context.Background(), transport.CreateTopicRequest{Name: req.TopicName})
	case GetTopicByIDOperation:
		return topicEndpoints.GetTopicByID(context.Background(), transport.GetTopicByIDRequest{ID: req.TopicID})
	case CreateRoomOperation:
		return roomEndpoints.CreateRoom(context.Background(), transport2.CreateRoomRequest{TopicID: req.TopicID})
	case DeleteRoomOperation:
		return roomEndpoints.DeleteRoom(context.Background(), transport2.DeleteRoomRequest{RoomID: req.RoomID})
	case InsertMessageOnRoomOperation:
		return roomEndpoints.InsertMessageOnRoom(context.Background(), transport2.InsertMessageOnRoomRequest{
			TopicID: req.TopicID,
			RoomID:  req.RoomID,
			Payload: req.Payload,
		})
	case GetMessageFromRoomOperation:
		return roomEndpoints.GetMessageFromRoom(context.Background(), transport2.GetMessageFromRoomRequest{
			RoomID:  req.RoomID,
			Index: req.Index,
		})
	}
	return resp, errors.New(fmt.Sprintf("invalid operation %s", req.Operation))
}

func (s server) Send(conn net.Conn, payload []byte) error {
	_, err := conn.Write(payload)
	return err
}
