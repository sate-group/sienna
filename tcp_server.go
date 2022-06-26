package sienna

import "net"

const (
	TCP_SERVER_NETWORK = "tcp"
)

type TcpServer struct {
	l       net.Listener
	address string
}

func newTcpServer(address string) (*TcpServer, error) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	server := &TcpServer{
		l:       l,
		address: address,
	}
	return server, nil
}

func (s *TcpServer) Listener() net.Listener { return s.l }

func (s *TcpServer) Address() string { return s.address }

func (s *TcpServer) Network() string { return "tcp" }

func (s *TcpServer) Accept() (Client, error) {
	l := s.l
	conn, err := l.Accept()
	if err != nil {
		return nil, err
	}
	var client Client = &TcpClient{
		conn: conn,
	}
	return client, nil
}
func (s *TcpServer) Close() error {
	l := s.Listener()
	err := l.Close()
	return err
}
