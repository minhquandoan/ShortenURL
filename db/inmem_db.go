package db

type MemDb struct {
	Store map[string]string
}

func NewMemDb() *MemDb {
	return &MemDb{
		Store: make(map[string]string),
	}
}