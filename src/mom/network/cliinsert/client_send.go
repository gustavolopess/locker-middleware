package main

import (
	"github.com/gustavolopess/locker/src/mom/network"
	logger2 "github.com/gustavolopess/locker/src/utils/logger"
)

func main() {
	logger := logger2.BuildLogger("client")
	client := network.NewClient("localhost", 8089, logger)
	defer client.Close()

	err := client.InsertMessageOnRoom("90859d15-661d-47cd-8e24-bd7c0423d42b", "DALE")
	if err != nil {
		panic(err)
	}
}