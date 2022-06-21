package engine

import (
	"fmt"
	"net"

	"github.com/google/uuid"
)

type ClientOptions struct {
	Address string
}

type Client struct {
	Id      uuid.UUID
	Conn    net.Conn
	Address string
}

func NewClient(opts *ClientOptions) *Client {
	client := &Client{
		Address: opts.Address,
	}
	return client
}

func (c *Client) Dial() error {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		return err
	}

	c.Conn = conn
	return nil
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) Send(a ...any) error {
	s := fmt.Sprint(a...)
	b := []byte(s)
	b = append(b, DIVIDER)

	_, err := c.Conn.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Sendf(format string, a ...any) error {
	s := fmt.Sprintf(format, a...)
	return c.Send(s)
}
