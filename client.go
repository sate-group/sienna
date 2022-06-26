package sienna

import "net"

const (
	DIVIDER = '\n'
)

type UnknownClientKindError string

func (e UnknownClientKindError) Error() string { return "Unknown client kind " + string(e) }

type Client interface {
	Conn() net.Conn
	Address() string
	Network() string

	Close() error

	Send(a ...any) (bool, error)
	Read() (string, error)
	SendJson(v any) error
	ReadJson(v any) error
}

func NewClient(kind, address string) (Client, error) {
	switch kind {
	case "tcp":
		client, err := newTcpClient(address)
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, UnknownClientKindError(kind)
	}
}
