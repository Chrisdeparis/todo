package server

import (
	"log"
	"net/http"
	"todo/controller"
	"todo/handler"
	"todo/model"
)

var routes = model.Routes{}

func Listen(addr string) {
	err := http.ListenAndServe(addr, getRouter())
	if err != nil {
		log.Fatal(err)
	}
}

func getRouter() http.Handler {
	indexController := controller.NewIndex()
	routes := append(routes, indexController.Routes...)

	//taskController := controller.NewTask()
	//routes = append(routes, taskController.Routes...)

	return handler.NewRouter(routes)
}
