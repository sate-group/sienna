package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/google/uuid"
)

const (
	DIVIDER = '\n'
)

type ClientOptions struct {
	Address string
}

type Client struct {
	Id   uuid.UUID
	Conn net.Conn
}

func NewClient(opts *ClientOptions) (*Client, error) {
	conn, err := net.Dial("tcp", opts.Address)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Conn: conn,
	}
	return client, nil
}

func (c *Client) Close() error {
	conn := c.Conn
	err := conn.Close()
	return err
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

func (c *Client) SendStruct(v any) error {
	out, err := json.Marshal(v)
	if err != nil {
		return err
	}
	str := string(out)
	if err := c.Sendf(str); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReadString() (string, error) {
	conn := c.Conn
	input, err := bufio.NewReader(conn).ReadString(DIVIDER)
	if err != nil {
		return "", err
	}
	str := strings.ReplaceAll(input, string(DIVIDER), "")
	return str, nil
}

func (c *Client) ReadStruct(v any) error {
	str, err := c.ReadString()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return err
	}
	return nil
}
