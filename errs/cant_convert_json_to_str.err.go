package errs

type CantConvertJsonToStrErr struct {
}

func IsCantConvertJsonToStrErr(err error) bool {
	switch err.(type) {
	case *CantConvertJsonToStrErr:
		return true
	}
	return false
}

func NewCantConvertJsonToStrErr() error {
	err := &CantConvertJsonToStrErr{}
	return err
}

func (e *CantConvertJsonToStrErr) Error() string {
	return "can't convert from JSON object to string."
}
