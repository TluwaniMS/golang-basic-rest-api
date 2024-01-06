package auxiliaries

import (
	"encoding/json"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func GenerateResponseMessage(message string) (responseMessage ResponseMessage){
	responseMessage,error := json.Marshal(message)

	return 
}