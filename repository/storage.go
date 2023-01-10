package repository

type Storage interface {
	Get(value interface{}) (interface{}, error)
}

type ServiceDB struct {
	MyDB Storage
}

func NewServiceDb(db Storage) *ServiceDB {
	return &ServiceDB{
		MyDB: db,
	}
}
