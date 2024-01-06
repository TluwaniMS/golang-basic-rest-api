package first_controller

import (
	"net/http"
	"encoding/json"
)

func GetHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
}

func PostHandler(response http.ResponseWriter, request *http.Request) {

}

func PutHandler(response http.ResponseWriter, request *http.Request) {

}

func PostHandler(response http.ResponseWriter, request *http.Request) {

}
