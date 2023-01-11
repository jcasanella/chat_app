package repository

type Storage interface {
	Get(value string) (string, error)
}

type ServiceDB struct {
	MyDB Storage
}

func NewServiceDb(db Storage) *ServiceDB {
	return &ServiceDB{
		MyDB: db,
	}
}
