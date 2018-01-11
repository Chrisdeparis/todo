package store

type StoreType int

const (
	MockType StoreType = iota
	MongoType
)

func CreateManager(t StoreType) (Manager, error) {
	switch t {
	case MongoType:
		return nil, nil
	default:
		return NewMockManager(), nil
	}
}
