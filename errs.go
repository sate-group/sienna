package sienna

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string { return "unknown network " + string(e) }
