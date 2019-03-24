package main

import (
	"time"

	"github.com/joyciapp/joyci-grpc/grpc/api"
)

func main() {
	go api.Serve()
	api.GitClone("git@github.com:joyciapp/joyci-grpc.git")
	time.Sleep(30 * time.Second)
}
