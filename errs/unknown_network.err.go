package errs

import "fmt"

type UnknownNetworkErr struct {
	network string
}

func IsUnknownNetworkErr(err error) bool {
	switch err.(type) {
	case *UnknownNetworkErr:
		return true
	}
	return false
}

func NewUnknownNetworkErr(network string) error {
	err := &UnknownNetworkErr{
		network: network,
	}
	return err
}

func (e *UnknownNetworkErr) Error() string {
	return fmt.Sprintf("unknown network. %s", e.network)
}
