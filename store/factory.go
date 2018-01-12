package store

import (
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type StoreType int

const (
	MockType StoreType = iota
	MongoType
)

func CreateManager(t StoreType) (Manager, error) {
	switch t {
	case MongoType:
		dsn := os.Getenv("MONGO_DSN")
		timeout := 5 * time.Second

		session, err := mgo.DialWithTimeout(dsn, timeout)
		if err != nil {
			return nil, err
		}

		session.SetSocketTimeout(timeout)
		session.SetPoolLimit(30)

		return NewMongoManager(session), nil
	default:
		return NewMockManager(), nil
	}
}
