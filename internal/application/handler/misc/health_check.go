package misc

import (
	"bookstoreapi/internal/application/writer"
	"log"
	"net/http"
)

// Returns a "Hello, World" message, to check if server is on
func HealthCheck(response http.ResponseWriter, request *http.Request) {
	body, err := writer.ToJSON("Hello, World!")

	if err != nil {
		log.Fatalf("error during healthcheck: %s", err)
	}

	response.Write([]byte(body))
}
