package main

import (
	"errors"
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
	client, err := sienna.NewClient("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	log.Print("client has successfully connected to the server")

	client.SendErr(errors.New("can't load user info."))
	log.Print("success send event")
}
