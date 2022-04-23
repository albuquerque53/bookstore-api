package misc

import (
	"bookstore_api/app/writer"
	"net/http"
)

// Returns a "Hello, World" message, to check if server is on
func HealthCheck(response http.ResponseWriter, request *http.Request) {
	jsonWriter := writer.NewMessageWriter("Hello, World!")

	reponseBody, _ := jsonWriter.JSONString()

	response.Write([]byte(reponseBody))
}
