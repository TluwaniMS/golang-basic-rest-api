package second_controller

import (
	"net/http"
	"basic-go-rest-api/auxiliaries"
	"basic-go-rest-api/auxiliary_types"
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

	body,_ = io.ReadAll(request.body)

	person auxiliary_types.Person

	error := json.Unmarshal(body,&person)

	message := auxiliaries.GenerateResponseMessage("Hello "+ person.name + " " + person.surname + " " + ", this is a response from the second controller.")
	
	response.Write(message)
}

func PutHandler(response http.ResponseWriter,request *http.Request) {

}

func DeleteHandler(response http.ResponseWriter,request *http.Request) {

}
