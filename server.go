package sienna

import (
	"net"

	"github.com/sate-infra/sienna/errs"
)

type Server interface {
	Listener() net.Listener
	Address() string
	Network() string

	Accept() (Client, error)
	Close() error
}

func NewServer(network, address string) (Server, error) {
	switch network {
	case TCP_SERVER_NETWORK:
		server, err := newTcpServer(address)
		if err != nil {
			return nil, err
		}
		return server, nil
	default:
		return nil, errs.NewUnknownNetworkErr(network)
	}
}
