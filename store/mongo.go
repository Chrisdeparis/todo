package store

import (
	"time"
	"todo/model"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

const collection = "tasks"

type mongo struct {
	session *mgo.Session
}

func NewMongoManager(s *mgo.Session) Manager {
	return &mongo{session: s}
}

func (m *mongo) Create(t *model.Task) error {
	session := m.session.Copy()
	defer session.Close()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	t.ID = id.String()
	t.CreatedAt = time.Now()

	return session.DB("").C(collection).Insert(t)
}

func (m *mongo) Find(ID string) (*model.Task, error) {
	if _, err := uuid.FromString(ID); err != nil {
		return nil, err
	}

	session := m.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)

	var task *model.Task
	err := c.FindId(ID).One(task)

	return task, err
}

func (m *mongo) Delete(ID string) error {
	if _, err := uuid.FromString(ID); err != nil {
		return err
	}

	session := m.session.Copy()
	defer session.Close()

	return session.DB("").C(collection).RemoveId(ID)
}

func (m *mongo) Update(t *model.Task) error {
	session := m.session.Copy()
	defer session.Close()

	return session.DB("").C(collection).UpdateId(t.ID, t)
}

func (m *mongo) All(offset, limit int) ([]model.Task, error) {
	return m.findAllBy(nil).Paginate(offset, limit)
}

func (m *mongo) findAllBy(query interface{}) *pagination {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("").C(collection)

	return &pagination{query: c.Find(query)}
}

type pagination struct {
	query *mgo.Query
}

func (p *pagination) Paginate(offset, limit int) (tasks []model.Task, err error) {
	tasks = []model.Task{}

	if offset > NoPaging && limit > NoPaging && limit > offset {
		err = p.query.Skip(offset).Limit(limit - offset).All(&tasks)
	} else {
		err = p.query.All(&tasks)
	}

	return
}
