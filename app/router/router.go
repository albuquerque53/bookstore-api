package router

import (
	"bookstore_api/app/controller/misc"
	"log"
	"net/http"
)

// Must handle all trafic of API
func HandleRequests() {
	http.HandleFunc("/health", misc.HealthCheck)

	log.Fatal(http.ListenAndServe(":2001", nil))
}
