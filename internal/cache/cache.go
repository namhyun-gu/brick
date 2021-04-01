package cache

type Cache interface {
	Write(content []byte) error

	Read() ([]byte, error)

	Exist() bool
}
