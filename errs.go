package sienna

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string { return "unknown network " + string(e) }

type SendDataFailedError string

func (e SendDataFailedError) Error() string { return "data transter failed" + string(e) }
