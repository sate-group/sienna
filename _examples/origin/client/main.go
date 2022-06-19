package main

import (
	"log"

	"github.com/sate-infra/sienna/engine"
)

func main() {
	opts := &engine.ClientOptions{
		Address: "localhost:9192",
	}
	client := engine.NewClient(opts)

	err := client.Dial()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err2 := client.Send("hello world")
	if err2 != nil {
		log.Println(err)
	}
}
