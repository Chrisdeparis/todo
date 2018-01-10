package controller

import (
	"fmt"
	"net/http"
	"todo/model"
)

type Index struct {
	Routes model.Routes
	Prefix string
}

func NewIndex() *Index {
	index := &Index{
		Routes: model.Routes{},
		Prefix: "/",
	}

	index.Routes = append(index.Routes, model.Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     index.Prefix,
		HandlerFunc: index.Index,
	})

	return index
}

func (i Index) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
