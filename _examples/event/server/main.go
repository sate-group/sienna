package main

import (
	"io"
	"log"

	"github.com/sate-infra/sienna"
)

type UserDto struct {
	Gender    string
	Race      string
	Birthday  string
	Street    string
	Telephone string
}

func main() {
	address := ":9192"
	server, err := sienna.NewServer("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server listening on %s", address)
	client, err := server.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	event, state, err := client.ReadEvent()
	if err == io.EOF {
		log.Fatal("no more input is available.")
	} else if err != nil {
		log.Fatal(err)
	}
	switch event {
	case "get_user":
		userDto := &UserDto{}
		if err := state.Decode(userDto); err != nil {
			log.Fatal(err)
		}
		log.Printf("client sent the dto %+v", userDto)
	default:
		log.Fatalf("unknown event name %s", event)
	}
}
