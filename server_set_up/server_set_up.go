package server_set_up

import (
	"basic-go-rest-api/first_controller"
	"basic-go-rest-api/second_controller"
	"net/http"
	"github.com/gorilla/mux"
)

func StartServer() {
	router := ConfigureRoutes()

	http.ListenAndServe(":3000", router)
}

func ConfigureRoutes() *mux.Router {
	router := mux.NewRouter()

	firstController := router.PathPrefix("/api/first-controller").Subrouter()
	secondController := router.PathPrefix("/api/second-controller").Subrouter()

	firstController.HandleFunc("", first_controller.GetHandler).Methods("GET")
	secondController.HandleFunc("", second_controller.GetHandler).Methods("GET")

	return router
}
