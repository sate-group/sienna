package sienna

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"syscall"

	"github.com/sate-infra/sienna/dtos"
	"github.com/sate-infra/sienna/errs"
)

const (
	TCP_CLIENT_NETWORK = "tcp"
)

var ()

type TcpClient struct {
	conn    net.Conn
	address string
	input   chan string
}

func newTcpClient(address string) (*TcpClient, error) {
	conn, err := net.Dial(TCP_CLIENT_NETWORK, address)
	if err != nil {
		return nil, err
	}
	client := &TcpClient{
		conn:    conn,
		address: address,
		input:   make(chan string),
	}
	return client, nil
}

func (c *TcpClient) Conn() net.Conn  { return c.conn }
func (c *TcpClient) Address() string { return c.address }
func (c *TcpClient) Network() string { return TCP_CLIENT_NETWORK }

func (c *TcpClient) Close() error {
	conn := c.Conn()
	close(c.input)
	if err := conn.Close(); err != nil {
		return err
	}
	return nil
}

func (c *TcpClient) Run() error {
	conn := c.Conn()
	str, err := bufio.NewReader(conn).ReadString(DIVIDER)
	if err == io.EOF {
		return errs.NewClientDisconnectedErr()
	} else if err != nil {
		return err
	}
	result := strings.ReplaceAll(str, string(DIVIDER), "")
	c.input <- result
	return nil
}

func (c *TcpClient) Send(a ...any) error {
	s := fmt.Sprint(a...)
	b := []byte(s)
	b = append(b, DIVIDER)

	n, err := c.Conn().Write(b)
	if errors.Is(err, syscall.EPIPE) {
		return errs.NewClientDisconnectedErr()
	} else if err != nil {
		return err
	}
	if n != len(b) {
		return errs.NewDataTransferFailedErr()
	}
	return nil
}

func (c *TcpClient) Read() (string, error) {
	str, ok := <-c.input
	if !ok {
		return "", errs.NewClientClosedErr()
	}
	return str, nil
}

func (c *TcpClient) SendJson(v any) error {
	data, err := jsonToStr(v)
	if err != nil {
		return err
	}

	if err := c.Send(data); err != nil {
		return err
	}
	return nil
}

func (c *TcpClient) ReadJson(v any) error {
	str, err := c.Read()
	if err != nil {
		return err
	}
	if err := strToJson(v, str); err != nil {
		return err
	}
	return nil
}

func (c *TcpClient) SendEvent(name string, v any) error {
	payload, err := jsonToStr(v)
	if err != nil {
		return err
	}
	dto := &dtos.EventDto{
		Name:    name,
		Payload: payload,
	}
	if err := c.SendJson(dto); err != nil {
		return err
	}
	return nil
}

func (c *TcpClient) ReadEvent() (string, *State, error) {
	dto := &dtos.EventDto{}
	if err := c.ReadJson(dto); err != nil {
		return "", nil, err
	}
	name := dto.Name
	p := dto.Payload
	s := NewState(p)
	return name, s, nil
}
