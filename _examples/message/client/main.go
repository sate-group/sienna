package main

import (
	"log"

	engine "github.com/sate-infra/sienna"
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
	log.Printf("Client has successfully connected to the server")

	client.Send("hello!")
}
