package main

import (
	"io"

	"github.com/sate-infra/sienna/engine"
	"github.com/sate-infra/sienna/logger"
)

func main() {
	opts := &engine.ServerOptions{
		Port: 9192,
	}
	server, err := engine.NewServer(opts)
	if err != nil {
		panic(err)
	}
	logger.Infof("Server listening on port %d", opts.Port)
	for {
		client, err := server.Accept()
		if err != nil {
			logger.Warnf(err.Error())
			continue
		}
		go handleClient(client)
	}
}

func handleClient(client *engine.Client) {
	defer client.Close()
	for {
		str, err := client.ReadString()
		if err == io.EOF {
			logger.Warnf("No more input is available.")
			return
		} else if err != nil {
			logger.Warnf(err.Error())
			return
		}
		logger.Infof("Client sent the message \"%s\"", str)
	}
}
