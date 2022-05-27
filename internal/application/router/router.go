package router

import (
	"bookstoreapi/internal/application/handler/author"
	"bookstoreapi/internal/application/handler/misc"
	"log"
	"net/http"
)

// Must handle all trafic of API
func HandleRequests() {
	http.HandleFunc("/health", misc.HealthCheck)
	http.HandleFunc("/authors/list", author.ListAuthors)
	http.HandleFunc("/authors/get/", author.GetAuthor)
	http.HandleFunc("/authors/new", author.NewAuthor)

	log.Fatal(http.ListenAndServe(":2001", nil))
}
