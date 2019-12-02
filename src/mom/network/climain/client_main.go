package main

import (
	"fmt"
	"github.com/gustavolopess/locker/src/mom/network"
	"github.com/gustavolopess/locker/src/mom/room_service/transport"
	logger2 "github.com/gustavolopess/locker/src/utils/logger"
)

func main() {
	logger := logger2.BuildLogger("client")
	client := network.NewClient("localhost", 8089, logger)
	defer client.Close()

	client.Subscribe("90859d15-661d-47cd-8e24-bd7c0423d42b", func(resp transport.GetMessageFromRoomResponse) {
		fmt.Println(resp.Message)
	})

	//err := client.Send(
	//	[]byte(fmt.Sprintf("{\"operation\": \"%s\", \"room_id\": \"90859d15-661d-47cd-8e24-bd7c0423d42b\" }\n",
	//	network.GetMessageFromRoomOperation)))
	//if err != nil {
	//	panic(err)
	//}
	//
	//ch := make(chan []byte)
	//
	//go func() {
	//	responseBytes, err := client.Receive()
	//	if err != nil {
	//		_ = level.Debug(logger).Log("client-listen", err)
	//		return
	//	}
	//
	//	_ = level.Info(logger).Log("client-listen", string(responseBytes))
	//
	//	ch <- responseBytes
	//}()
	//
	//fmt.Println(<- ch)
}
