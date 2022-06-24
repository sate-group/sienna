package sienna

import "net"

type UnknownServerKindError string

func (e UnknownServerKindError) Error() string { return "Unknown server kind " + string(e) }

type Server interface {
	Listener() net.Listener
	Address() string
	Kind() string

	Accept() (Client, error)
	Close() error
}

func NewServer(kind string, address string) (Server, error) {
	switch kind {
	case "tcp":
		server, err := newTcpServer(address)
		if err != nil {
			return nil, err
		}
		return server, nil
	default:
		return nil, UnknownServerKindError(kind)
	}
}
