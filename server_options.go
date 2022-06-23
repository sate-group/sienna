package engine

const (
	DEFULAT_PORT = uint16(4000)
)

type ServerOptions struct {
	Port uint16
}

func (opts *ServerOptions) getPort() uint16 {
	if opts != nil && opts.Port != 0 {
		return opts.Port
	}
	return DEFULAT_PORT
}
