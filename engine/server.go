package engine

import (
	"net"
	"strconv"
)

type Server struct {
	Listener net.Listener
	Port     uint16
}

func NewServer(options *ServerOptions) (*Server, error) {
	address := ":" + strconv.Itoa(int(options.Port))
	l, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	server := &Server{
		Listener: l,
		Port:     options.getPort(),
	}

	return server, nil
}

func (s *Server) Close() error {
	l := s.Listener
	err := l.Close()
	return err
}

func (s *Server) Accept() (*Client, error) {
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
