package main

import (
	"github.com/gustavolopess/locker/src/mom/network"
	logger2 "github.com/gustavolopess/locker/src/utils/logger"
)

func main() {
	logger := logger2.BuildLogger("topic")
	client := network.NewClient("localhost", 8089, logger)
	defer client.Close()

	err := client.CreateTopic("BuildingTopic")
	if err != nil {
		panic(err)
	}
}