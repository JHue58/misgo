package request

type Code int64

const (
	Success Code = iota
	ParmaError
	DataBaseError
	ServerError
)
