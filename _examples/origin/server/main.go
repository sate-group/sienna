package main

import (
	"github.com/sate-infra/sienna/engine"
	"github.com/sate-infra/sienna/logger"
)

func main() {
	opts := &engine.ServerOptions{
		Port: 9192,
	}
	server := engine.NewServer(opts)

	server.OnString(func(client *engine.Client, data string) {
		logger.Infof("Client %s sent the message \"%s\"\n", client.Id.String()[1:8], data)

	})

	logger.Infof("\033[32mServer listening on port %d\033[0m", opts.Port)
	err2 := server.Listen()
	if err2 != nil {
		panic(err2)
	}
}
