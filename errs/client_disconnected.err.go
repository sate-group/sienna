package errs

type ClientDisconnectedErr struct {
}

func IsClientDisconnectedErr(err error) bool {
	switch err.(type) {
	case *ClientDisconnectedErr:
		return true
	}
	return false
}

func NewClientDisconnectedErr() error {
	err := &ClientDisconnectedErr{}
	return err
}

func (e *ClientDisconnectedErr) Error() string {
	return "client has disconnected."
}
