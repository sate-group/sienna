package errs

type DataTransferFailedErr struct {
}

func IsDataTransferFailedErr(err error) bool {
	switch err.(type) {
	case *DataTransferFailedErr:
		return true
	}
	return false
}

func NewDataTransferFailedErr() error {
	err := &DataTransferFailedErr{}
	return err
}

func (e *DataTransferFailedErr) Error() string {
	return "data transfer failed."
}
