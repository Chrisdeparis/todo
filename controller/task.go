package controller

import (
	"net/http"
	"todo/model"
	"todo/store"
)

type Task struct {
	store  store.Manager
	Routes model.Routes
	Prefix string
}

func NewTask(m store.Manager) *Task {
	prefix := "/todo/tasks"

	t := &Task{
		store:  m,
		Prefix: prefix,
	}

	t.Routes = append(t.Routes, model.Routes{
		model.Route{
			Name:        "All",
			Method:      "GET",
			Pattern:     prefix,
			HandlerFunc: t.All,
		},
		model.Route{
			Name:        "Create",
			Method:      "POST",
			Pattern:     prefix,
			HandlerFunc: t.Create,
		},
		model.Route{
			Name:        "Find",
			Method:      "GET",
			Pattern:     prefix + "/{id}",
			HandlerFunc: t.FindById,
		},
		model.Route{
			Name:        "Update",
			Method:      "PUT",
			Pattern:     prefix,
			HandlerFunc: t.Update,
		},
		model.Route{
			Name:        "Delete",
			Method:      "DELETE",
			Pattern:     prefix + "/{id}",
			HandlerFunc: t.Delete,
		},
	}...)

	return t
}

// All returns all tasks.
func (t *Task) All(w http.ResponseWriter, r *http.Request) {

}

// Create creates a new task.
func (t *Task) Create(w http.ResponseWriter, r *http.Request) {

}

// FindById find a task by its ID.
func (t *Task) FindById(w http.ResponseWriter, r *http.Request) {

}

func (t *Task) Update(w http.ResponseWriter, r *http.Request) {

}

func (t *Task) Delete(w http.ResponseWriter, r *http.Request) {

}
