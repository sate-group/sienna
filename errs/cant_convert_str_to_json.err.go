package errs

type CantConvertStrToJsonErr struct {
}

func IsCantConvertStrToJsonErr(err error) bool {
	switch err.(type) {
	case *CantConvertStrToJsonErr:
		return true
	}
	return false
}

func NewCantConvertStrToJsonErr() error {
	err := &CantConvertStrToJsonErr{}
	return err
}

func (e *CantConvertStrToJsonErr) Error() string {
	return "can't convert from string to JSON object."
}
