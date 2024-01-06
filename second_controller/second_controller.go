package second_controller

import (
	"net/http"
	"basic-go-rest-api/auxiliaries"
	"io"
)

func GetHandler(response http.ResponseWriter,request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	message := auxiliaries.GenerateResponseMessage("Response from Get Method on the Second Controller ;) ...!!")

	response.Write(message)
}

func PostHandler(response http.ResponseWriter,request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	body,error = io.ReadAll(request.body)

	if error != nil{

	}

	
	
}

func PutHandler(response http.ResponseWriter,request *http.Request) {

}

func DeleteHandler(response http.ResponseWriter,request *http.Request) {

}
