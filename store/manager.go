package store

import "todo/model"

const NoPaging = -1

type Manager interface {
	All(offset, limit int) ([]model.Task, error)

	Create(*model.Task) error

	Find(ID string) (*model.Task, error)

	Update(*model.Task) error

	Delete(ID string) error
}
