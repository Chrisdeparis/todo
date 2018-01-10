package store

import (
	"errors"
	"time"
	"todo/model"

	"github.com/satori/go.uuid"
)

type Mock struct {
	tasks map[string]*model.Task
}

func (m *Mock) Create(t *model.Task) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	t.ID = id.String()
	t.CreatedAt = time.Now()

	m.tasks[t.ID] = t
	return nil
}

func (m *Mock) All(offset, limit int) ([]model.Task, error) {
	tasks := []model.Task{}

	if offset == NoPaging {
		offset = 0
	}
	if limit == NoPaging {
		limit = len(m.tasks)
	}
	if offset > limit || limit > len(m.tasks) {
		return tasks, nil
	}

	for _, task := range m.tasks {
		tasks = append(tasks, *task)
	}

	return tasks[offset:limit], nil
}

func (m *Mock) Find(ID string) (*model.Task, error) {
	if task, ok := m.tasks[ID]; ok {
		return task, nil
	}
	return nil, errors.New("task not found")
}

func (m *Mock) Delete(ID string) error {
	delete(m.tasks, ID)
	return nil
}

func (m *Mock) Update(t *model.Task) error {
	if _, ok := m.tasks[t.ID]; ok {
		m.tasks[t.ID] = t
	}
	return nil
}
