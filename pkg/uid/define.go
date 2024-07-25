package uid

type Manager interface {
	Generate() (Uid string, err error)
}
