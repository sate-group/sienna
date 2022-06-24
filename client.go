package sienna

type Client interface {
	Close() error
	Send(a ...any) error
	SendStruct(v any) error
	ReadString() (string, error)
	ReadStruct(v any) error
}

type ClientOptions struct {
	Address string
}
