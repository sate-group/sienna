package sienna

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/sate-infra/sienna/dtos"
	"github.com/sate-infra/sienna/errs"
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
	data, err := jsonToStr(v)
	if err != nil {
		return err
	}
	dto := &dtos.SendJsonDto{
		Err:  "",
		Data: data,
	}
	str, err := jsonToStr(dto)
	if err != nil {
		return err
	}

	if ok, err := c.Send(str); err != nil {
		return err
	} else if !ok {
		return SendDataFailedError(str)
	}
	return nil
}

func jsonToStr(v any) (string, error) {
	out, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	str := string(out)
	return str, nil
}

func (c *TcpClient) ReadJson(v any) error {
	str, err := c.Read()
	if err != nil {
		return err
	}
	dto := &dtos.ReadJsonDto{}
	if err := json.Unmarshal([]byte(str), &dto); err != nil {
		return err
	}
	if dto.Err != "" {
		return errs.NewClientErr(dto.Err)
	}
	data := dto.Data
	if err := json.Unmarshal([]byte(data), &v); err != nil {
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

func (c *TcpClient) SendErr(err error) error {
	errStr := err.Error()
	dto := &dtos.SendJsonDto{
		Err:  errStr,
		Data: "",
	}
	str, err := jsonToStr(dto)
	if err != nil {
		return err
	}
	if ok, err := c.Send(str); err != nil {
		return err
	} else if !ok {
		return SendDataFailedError(str)
	}
	return nil
}
