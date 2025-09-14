package store


type Store interface {
	Write(b []byte, f string) error
	Read(f string) ([]byte, error)
	Exists(f string) (bool, error)
}

