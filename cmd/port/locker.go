package main

import (
	"flag"
	"fmt"
	"github.com/gustavolopess/locker/src/mom/network"
	logger2 "github.com/gustavolopess/locker/src/utils/logger"
)

func main() {
	lockerID := flag.String("id", "", "locker id")
	lockerOperation := flag.String("op", "", "operation to be executed")
	flag.Parse()

	if len(*lockerID) == 0 {
		panic("invalid locker id")
	}

	if *lockerOperation != "open" && *lockerOperation != "close" {
		panic("locker operations must be 'open' or 'close'")
	}

	logger := logger2.BuildLogger("client")
	client := network.NewClient("localhost", 8089, logger)
	defer client.Close()

	err := client.InsertMessageOnRoom(
		"ee235856-1238-4f16-a2a3-8b3389bd5e11",
		fmt.Sprintf("%s;%s", *lockerID, *lockerOperation),
	)
	if err != nil {
		panic(err)
	}
}
