package misc

import (
	"fmt"
	"net/http"
)

// Returns a "Hello, World" message, to check if server is on
func HealthCheck(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello, World!")
}
