package second_controller

import (
	"basic-go-rest-api/auxiliaries"
	"basic-go-rest-api/auxiliary_types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	message := auxiliaries.GenerateResponseMessage("Response from Get Method on the Second Controller ;) ...!!")

	response.Write(message)
}

func PostHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	body, _ := ioutil.ReadAll(request.Body)

	var person auxiliary_types.Person

	error := json.Unmarshal(body, &person)

	if error != nil {
		fmt.Println("There was an error.")
	}

	message := auxiliaries.GenerateResponseMessage("Hello " + person.Name + " " + person.Surname + " " + ", this is a response from the second controller.")

	response.Write(message)
}

func PutHandler(response http.ResponseWriter, request *http.Request) {

}

func DeleteHandler(response http.ResponseWriter, request *http.Request) {

}
