package network

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gustavolopess/locker/src/mom/room_service/transport"
	"net"
	"time"
)

type Client interface {
	Close() error
	Send([]byte) error
	Receive() ([]byte, error)
	Subscribe(roomID string, callback func(resp transport.GetMessageFromRoomResponse))
	InsertMessageOnRoom(roomID string, payload string) error
}

type client struct {
	ServerHost string
	ServerPort int
	Logger log.Logger
	conn net.Conn
}

func NewClient(serverHost string, serverPort int, logger log.Logger) Client {
	c :=  &client{
		ServerHost: serverHost,
		ServerPort: serverPort,
		Logger: logger,
	}

	if err := c.connect(); err != nil {
		panic(err)
	}

	return c
}

func (c *client) connect() error {
	var err error
	c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.ServerHost, c.ServerPort))
	return err
}

func (c *client) Close() error {
	if c.conn == nil {
		return errors.New("Connection is nil")
	}

	return c.conn.Close()
}

func (c *client) Send(msg []byte) error {
	_, err := c.conn.Write(msg)
	return err
}

func (c *client) Receive() ([]byte, error) {
	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	return []byte(message), nil
}

func (c *client) GetMessageFromRoom(roomID string, lastIndex int64) error {
	req := GenericRequest{
		Operation: GetMessageFromRoomOperation,
		RoomID:    roomID,
		Index: lastIndex,
	}

	reqBytes, err := json.Marshal(&req)
	reqBytes = append(reqBytes, '\n')

	if err != nil {
		_ = level.Debug(c.Logger).Log("client-receive", err)
		return err
	}

	return c.Send(reqBytes)
}

func (c *client) InsertMessageOnRoom(roomID string, payload string) error {
	req := GenericRequest{
		Operation: InsertMessageOnRoomOperation,
		RoomID:    roomID,
		Payload:   payload,
	}

	reqBytes, err := json.Marshal(&req)
	reqBytes = append(reqBytes, '\n')
	if err != nil {
		_ = level.Debug(c.Logger).Log("client-insert", err)
		return err
	}

	return c.Send(reqBytes)
}

func (c *client) Subscribe(roomID string, callback func(resp transport.GetMessageFromRoomResponse)) {
	var lastIdx int64 = -1
	for {
		err := c.GetMessageFromRoom(roomID, lastIdx)
		if err != nil {
			panic(err)
		}

		_ = level.Info(c.Logger).Log("client-subscribe", "asked for message")

		responseBytes, err := c.Receive()
		if err != nil {
			_ = level.Debug(c.Logger).Log("client-subscribe", err)
			return
		}

		_ = level.Info(c.Logger).Log("client-subscribe", string(responseBytes))

		var response transport.GetMessageFromRoomResponse
		if err = json.Unmarshal(responseBytes, &response); err != nil {
			panic(err)
		}

		if lastIdx < response.Message.Index {
			lastIdx = response.Message.Index
			go callback(response)
		}

		time.Sleep(1 * time.Second)
	}
}
