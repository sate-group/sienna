package main

import (
	"log"

	"github.com/sate-infra/sienna"
)

func main() {
	address := ":9192"
	client, err := sienna.NewClient("tcp", address)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	log.Printf("Client has successfully connected to the server")

	client.Send("hello!")
}
