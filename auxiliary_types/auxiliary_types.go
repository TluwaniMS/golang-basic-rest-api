package auxiliary_types

type ResponseMessage struct {
	Message string `json:"message"`
}

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
}
