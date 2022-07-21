package sienna

import "net"

const (
	DIVIDER = '\n'
)

type Client interface {
	Conn() net.Conn
	Address() string
	Network() string

	Close() error

	Send(a ...any) (bool, error)
	Read() (string, error)
	SendJson(v any) error
	ReadJson(v any) error
	SendEvent(name string, v any) error
	ReadEvent() (string, *State, error)
	SendErr(err error) error
}

func NewClient(network, address string) (Client, error) {
	switch network {
	case TCP_CLIENT_NETWORK:
		client, err := newTcpClient(address)
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, UnknownNetworkError(network)
	}
}
