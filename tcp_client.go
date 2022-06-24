package sienna

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type TcpClient struct {
	conn    net.Conn
	address string
}

func newTcpClient(address string) (*TcpClient, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	client := &TcpClient{
		conn:    conn,
		address: address,
	}
	return client, nil
}

func (c *TcpClient) Conn() net.Conn {
	return c.conn
}
func (c *TcpClient) Address() string {
	return c.address
}
func (c *TcpClient) Kind() string {
	return "tcp"
}

func (c *TcpClient) Close() error {
	conn := c.Conn()
	err := conn.Close()
	return err
}

func (c *TcpClient) Send(a ...any) (bool, error) {
	s := fmt.Sprint(a...)
	b := []byte(s)
	b = append(b, DIVIDER)

	n, err := c.Conn().Write(b)
	if err != nil {
		return false, err
	}
	success := false
	if n == len(b) {
		success = true
	}
	return success, nil // n - 1 is removed "DIVIDER"
}

func (c *TcpClient) Read() (string, error) {
	conn := c.Conn()
	str, err := bufio.NewReader(conn).ReadString(DIVIDER)
	if err != nil {
		return "", err
	}
	result := strings.ReplaceAll(str, string(DIVIDER), "")
	return result, nil
}
