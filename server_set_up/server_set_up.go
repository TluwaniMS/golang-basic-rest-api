package server_set_up

import (
	"basic-go-rest-api/first_controller"
	"basic-go-rest-api/second_controller"
	"github.com/gorilla/mux"
)

func StartServer() {

}

func ConfigureRoutes() {
	router := mux.NewRouter()

	firstController := router.PathPrefix("/api/first-controller").Subrouter()
	secondController := router.PathPrefix("/api/second-controller").Subrouter()

}
