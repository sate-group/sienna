package sienna

type TcpServer struct {
	address string
}

func newTcpServer(address string) *TcpServer {
	return &TcpServer{
		address: address,
	}
}

func (s *TcpServer) Address() string {
	return s.address
}

func (s *TcpServer) Kind() string {
	return "tcp"
}

func (s *TcpServer) Accept() (Client, error) {
	return nil, nil
}
func (s *TcpServer) Close() error {
	return nil
}
