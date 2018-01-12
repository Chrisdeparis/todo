package server

import (
	"log"
	"net/http"
	"todo/controller"
	"todo/handler"
	"todo/model"
	"todo/store"
)

var routes = model.Routes{}

func Listen(addr string) {
	err := http.ListenAndServe(addr, getRouter())
	if err != nil {
		log.Fatal(err)
	}
}

func ListenTLS(addr string) {
	sslPath := "etc/ssl"
	err := http.ListenAndServeTLS(addr, sslPath+"/server.crt", sslPath+"/server.key", getRouter())
	if err != nil {
		log.Fatal(err)
	}
}

func getRouter() http.Handler {
	indexController := controller.NewIndex()
	routes := append(routes, indexController.Routes...)

	m, err := store.CreateManager(store.MongoType)
	if err != nil {
		log.Fatal(err)
	}

	taskController := controller.NewTask(m)
	routes = append(routes, taskController.Routes...)

	return handler.NewRouter(routes)
}
