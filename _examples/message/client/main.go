package main

import (
	"github.com/sate-infra/sienna/engine"
	"github.com/sate-infra/sienna/logger"
)

func main() {
	opts := &engine.ClientOptions{
		Address: "localhost:9192",
	}
	client, err := engine.NewClient(opts)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	logger.Infof("Client has successfully connected to the server")

	client.Send("hello!")
}
