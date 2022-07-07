package sienna

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

const (
	TCP_CLIENT_NETWORK = "tcp"
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

func (c *TcpClient) Conn() net.Conn  { return c.conn }
func (c *TcpClient) Address() string { return c.address }
func (c *TcpClient) Network() string { return "tcp" }

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

func (c *TcpClient) SendJson(v any) error {
	out, err := json.Marshal(v)
	if err != nil {
		return err
	}
	str := string(out)
	if ok, err := c.Send(str); err != nil {
		return err
	} else if !ok {
		return SendDataFailedError(str)
	}
	return nil
}

func (c *TcpClient) ReadJson(v any) error {
	str, err := c.Read()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return err
	}
	return nil
}

type EventDto struct {
	Name    string
	Payload string
}

func (c *TcpClient) SendEvent(name string, v any) error {
	out, err := json.Marshal(v)
	if err != nil {
		return err
	}
	payload := string(out)
	dto := &EventDto{
		Name:    name,
		Payload: payload,
	}
	if err := c.SendJson(dto); err != nil {
		return err
	}
	return nil
}
func (c *TcpClient) ReadEvent() (string, *State, error) {
	dto := &EventDto{}
	err := c.ReadJson(dto)
	if err != nil {
		return "", nil, err
	}
	name := dto.Name
	p := dto.Payload
	s := NewState(p)
	return name, s, nil
}
