package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/gustavolopess/locker/src/locker/store"
	"github.com/gustavolopess/locker/src/locker_service"
	transport2 "github.com/gustavolopess/locker/src/locker_service/transport"
	"github.com/gustavolopess/locker/src/mom/network"
	"github.com/gustavolopess/locker/src/mom/room_service/transport"
	logger2 "github.com/gustavolopess/locker/src/utils/logger"
	"strings"
)

func main() {
	logger := logger2.BuildLogger("client")
	client := network.NewClient("localhost", 8089, logger)
	defer client.Close()

	lockerStore := store.NewLockerStore(logger)
	lockerService := locker_service.NewLockerService(lockerStore, logger)
	lockerEndpoints := transport2.MakeLockerEndpoints(lockerService)

	client.Subscribe("ee235856-1238-4f16-a2a3-8b3389bd5e11", func(resp transport.GetMessageFromRoomResponse) {
		msg := resp.Message.Payload

		splitRes := strings.Split(msg, ";")
		lockerID, operation := splitRes[0], splitRes[1]

		var err error
		switch operation {
		case "open":
			_, err = lockerEndpoints.OpenLocker(context.Background(), transport2.OpenLockerRequest{lockerID})
			break
		case "close":
			_, err = lockerEndpoints.CloseLocker(context.Background(), transport2.CloseLockerRequest{lockerID})
			break
		}

		if err != nil {
			_ = level.Debug(logger).
				Log(
					"subscriber",
					fmt.Sprintf("%s @ %s FAILED", operation, lockerID),
				)
		} else {
			_ = level.Info(logger).
				Log(
					"subscriber",
					fmt.Sprintf("%s @ %s processed", operation, lockerID),
				)
		}
	})
}


