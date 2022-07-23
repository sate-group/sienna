package main

import (
	"log"

	"github.com/sate-infra/sienna"
	"github.com/sate-infra/sienna/errs"
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
		c, err := server.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go func() {
			defer c.Close()
			for {
				if err := c.Run(); errs.IsClientDisconnectedErr(err) {
					log.Print(err)
					return
				} else if err != nil {
					log.Print(err)
					return
				}
			}
		}()
		go func() {
			for {
				if err := handleClient(c); errs.IsClientClosedErr(err) {
					log.Print(err)
					return
				} else if errs.IsClientDisconnectedErr(err) {
					log.Print(err)
					return
				} else if errs.IsDataTransferFailedErr(err) {
					log.Print(err)
				} else if err != nil {
					log.Print(err)
				}
			}
		}()
	}
}

func handleClient(c sienna.Client) error {
	userDto := &UserDto{}
	if err := c.ReadJson(userDto); err != nil {
		return err
	}
	log.Printf("client sent the dto %+v", userDto)
	return nil
}
