package main

import (
	"io"
	"log"

	engine "github.com/sate-infra/sienna"
)

type UserDto struct {
	Gender    string `json:"gender"`
	Race      string `json:"race"`
	Birthday  string `json:"birthday"`
	Street    string `json:"street"`
	Telephone string `json:"telephone"`
}

func main() {
	opts := &engine.ServerOptions{
		Port: 9192,
	}
	server, err := engine.NewServer(opts)
	if err != nil {
		panic(err)
	}
	log.Printf("Server listening on port %d", opts.Port)
	for {
		client, err := server.Accept()
		if err != nil {
			log.Print(err.Error())
			continue
		}
		go handleClient(client)
	}
}

func handleClient(client *engine.Client) {
	defer client.Close()
	userDto := &UserDto{}
	err := client.ReadStruct(userDto)
	if err == io.EOF {
		log.Print("No more input is available.")
		return
	} else if err != nil {
		log.Print(err.Error())
		return
	}
	log.Printf("Client sent the dto %+v", userDto)
}
