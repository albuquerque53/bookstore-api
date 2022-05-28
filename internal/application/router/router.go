package router

import (
	"bookstoreapi/internal/application/handler/author"
	"bookstoreapi/internal/application/handler/misc"
	"bookstoreapi/internal/application/middleware"
	"log"
	"net/http"
)

// Must handle all trafic of API
func HandleRequests() {
	route("/health", misc.HealthCheck)
	route("/authors/list", author.ListAuthors)
	route("/authors/get/", author.GetAuthor)
	route("/authors/new", author.NewAuthor)
	route("/authors/delete/", author.DeleteAuthor)

	log.Fatal(http.ListenAndServe(":2001", nil))
}

func route(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, middleware.PanicHandler(handler))
}
