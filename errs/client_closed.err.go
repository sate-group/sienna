package errs

type ClientClosedErr struct {
}

func IsClientClosedErr(err error) bool {
	switch err.(type) {
	case *ClientClosedErr:
		return true
	}
	return false
}

func NewClientClosedErr() error {
	err := &ClientClosedErr{}
	return err
}

func (e *ClientClosedErr) Error() string {
	return "client closed."
}
