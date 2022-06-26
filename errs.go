package sienna

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string { return "Unknown network " + string(e) }

type SendDataFailedError string

func (e SendDataFailedError) Error() string { return "Data transter failed" + string(e) }
