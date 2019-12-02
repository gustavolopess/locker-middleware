package main

import (
	"context"
	"github.com/gustavolopess/locker/src/locker/store"
	"github.com/gustavolopess/locker/src/locker_service"
	"github.com/gustavolopess/locker/src/utils/logger"
)

func main() {
	logger := logger.BuildLogger("locker-creator")
	lockerStore := store.NewLockerStore(logger)
	svc := locker_service.NewLockerService(lockerStore, logger)

	_, err := svc.CreateLocker(context.Background(), "my-locker")
	if err != nil {
		panic(err)
	}
}
