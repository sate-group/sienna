package main

import (
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
	client.SendJson(user)
}
