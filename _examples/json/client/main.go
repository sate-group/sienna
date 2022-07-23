package main

import (
	"log"
	"time"

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
		panic(err)
	}
	defer client.Close()
	log.Print("Client has successfully connected to the server")

	user := &UserDto{
		Gender:    "male",
		Race:      "White",
		Birthday:  "2/26/1985 (37 years old)",
		Street:    "2551 Eastland Avenue",
		Telephone: "601-420-5622",
	}
	for {
		if err := client.SendJson(user); err != nil {
			log.Print(err)
		}
		log.Print("send user json to server")
		time.Sleep(1 * time.Second)
	}
}
