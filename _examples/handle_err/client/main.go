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
	client, err := sienna.NewClient("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	log.Print("client has successfully connected to the server")

	user := &UserDto{
		Gender:    "male",
		Race:      "White",
		Birthday:  "2/26/1985 (37 years old)",
		Street:    "2551 Eastland Avenue",
		Telephone: "601-420-5622",
	}
	/*
		Format/Syntax
			productName:objectName_actionName

		link: https://davidwells.io/blog/clean-analytics
	*/
	client.SendEvent("client:user_sendInfo", user)
	log.Print("success send event")
}
