package main

import (
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

func handleClient(c sienna.Client) {
	defer c.Close()
	userDto := &UserDto{}
	err := c.ReadJson(userDto)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("client sent the dto %+v", userDto)
}
