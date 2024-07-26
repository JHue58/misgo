package monitor

type UpdateFunc func() (ok bool, msg string, err error)

type Monitor interface {
	Start() error
	Stop()
}
