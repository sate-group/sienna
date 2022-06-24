package sienna

type UnknownServerKindError string

func (e UnknownServerKindError) Error() string { return "Unknown server kind " + string(e) }

type Server interface {
	Address() string
	Kind() string

	Accept() (Client, error)
	Close() error
}

func NewServer(kind string, address string) (Server, error) {
	var server Server

	switch kind {
	case "tcp":
		server = newTcpServer(address)
		return server, nil
	default:
		return nil, UnknownServerKindError(kind)
	}
}
