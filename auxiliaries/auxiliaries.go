package auxiliaries

import (
	"encoding/json"
	"basic-go-rest-api/auxiliary_types"
)


func GenerateResponseMessage(message string) (responseMessage auxiliary_types.ResponseMessage){
	responseMessage,error := json.Marshal(message)

	return 
}