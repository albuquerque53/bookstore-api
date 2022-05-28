package misc

import (
	"bookstoreapi/internal/application/writer"
	"net/http"
)

// Returns a "Hello, World" message, to check if server is on
func HealthCheck(response http.ResponseWriter, request *http.Request) {
	writer.SendResponse(response, 200, writer.JSONResponse{
		Message: "Hello, World",
		Data:    nil,
	})
}
