package errs

type ClientErr struct {
	str string
}

func IsClientErr(err error) bool {
	switch err.(type) {
	case *ClientErr:
		return true
	}
	return false
}

func NewClientErr(str string) error {
	err := &ClientErr{
		str: str,
	}
	return err
}

func (e *ClientErr) Error() string {
	return e.str
}
