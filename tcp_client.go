package sienna

import "net"

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
