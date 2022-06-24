package main

import (
	"io"
	"log"

	"github.com/sate-infra/sienna"
)

type UserDto struct {
	Gender    string `json:"gender"`
	Race      string `json:"race"`
	Birthday  string `json:"birthday"`
	Street    string `json:"street"`
	Telephone string `json:"telephone"`
}

func main() {
	address := ":9192"
	server, err := sienna.NewServer("tcp", address)
	if err != nil {
		panic(err)
	}
	log.Printf("Server listening on %s", address)
	for {
		client, err := server.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleClient(client)
	}
}

func handleClient(client sienna.Client) {
	defer client.Close()
	userDto := &UserDto{}
	err := client.ReadJson(userDto)
	if err == io.EOF {
		log.Print("No more input is available.")
		return
	} else if err != nil {
		log.Print(err.Error())
		return
	}
	log.Printf("Client sent the dto %+v", userDto)
}
