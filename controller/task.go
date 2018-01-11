package controller

import (
	"log"
	"net/http"
	"strconv"
	"todo/handler"
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
	var err error

	offset := store.NoPaging
	limit := store.NoPaging

	offset, err = strconv.Atoi(handler.GetParams("offset", r))
	if err != nil {
		offset = store.NoPaging
	}

	limit, err = strconv.Atoi(handler.GetParams("limit", r))
	if err != nil {
		limit = store.NoPaging
	}

	tasks, err := t.store.All(offset, limit)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "Error while retrieving tasks", http.StatusInternalServerError)
		return
	}

	handler.SendJSONOk(w, &tasks)
}

// Create creates a new task.
// http://localhost:7070/todo/tasks (POST)
func (t *Task) Create(w http.ResponseWriter, r *http.Request) {
	task := &model.Task{}

	err := handler.ParseJSON(r, task)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "Error while decoding to create", http.StatusBadRequest)
		return
	}

	err = t.store.Create(task)
	if err != nil {
		log.Println(err)
		handler.SendJSONError(w, "Error while creating task", http.StatusInternalServerError)
		return
	}

	handler.SendJSONWithStatus(w, task, http.StatusCreated)
}

// FindById find a task by its ID.
func (t *Task) FindById(w http.ResponseWriter, r *http.Request) {

}

func (t *Task) Update(w http.ResponseWriter, r *http.Request) {

}

func (t *Task) Delete(w http.ResponseWriter, r *http.Request) {

}
