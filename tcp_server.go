package sienna

import (
	"net"
	"strconv"
)

type TcpServer struct {
	Listener net.Listener
	Options  *ServerOptions
}

func NewTcpServer(opts *ServerOptions) (*Server, error) {
	port := opts.getPort()
	address := ":" + strconv.Itoa(int(port))
	l, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	var server Server = &TcpServer{
		Listener: l,
		Options:  opts,
	}

	return &server, nil
}

func (s *TcpServer) Close() error {
	l := s.Listener
	err := l.Close()
	return err
}

func (s *TcpServer) Accept() (*Client, error) {
	l := s.Listener
	conn, err := l.Accept()
	if err != nil {
		return nil, err
	}
	client := &Client{
		Conn: conn,
	}
	return client, nil
}
