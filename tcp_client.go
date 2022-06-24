package sienna

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

type TcpClient struct {
	Id   uuid.UUID
	Conn net.Conn
}

func NewTcpClient(opts *ClientOptions) (*Client, error) {
	conn, err := net.Dial("tcp", opts.Address)
	if err != nil {
		return nil, err
	}

	var client Client = &TcpClient{
		Conn: conn,
	}
	return &client, nil
}

func (c *TcpClient) Close() error {
	conn := c.Conn
	err := conn.Close()
	return err
}

func (c *TcpClient) Send(a ...any) error {
	s := fmt.Sprint(a...)
	b := []byte(s)
	b = append(b, DIVIDER)

	_, err := c.Conn.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (c *TcpClient) SendStruct(v any) error {
	out, err := json.Marshal(v)
	if err != nil {
		return err
	}
	str := string(out)
	if err := c.Send(str); err != nil {
		return err
	}
	return nil
}

func (c *TcpClient) ReadString() (string, error) {
	conn := c.Conn
	input, err := bufio.NewReader(conn).ReadString(DIVIDER)
	if err != nil {
		return "", err
	}
	str := strings.ReplaceAll(input, string(DIVIDER), "")
	return str, nil
}

func (c *TcpClient) ReadStruct(v any) error {
	str, err := c.ReadString()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return err
	}
	return nil
}
