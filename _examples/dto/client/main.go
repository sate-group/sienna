package main

import (
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
	opts := &engine.ClientOptions{
		Address: "localhost:9192",
	}
	client, err := engine.NewClient(opts)
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
	client.SendStruct(user)
}
