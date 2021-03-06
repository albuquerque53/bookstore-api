package router

import (
	"bookstoreapi/internal/application/handler/author"
	"bookstoreapi/internal/application/handler/book"
	"bookstoreapi/internal/application/handler/category"
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
	route("/authors/edit/", author.UpdateAuthor)
	route("/authors/delete/", author.DeleteAuthor)

	route("/categories/list", category.ListCategories)
	route("/categories/get/", category.GetCategory)
	route("/categories/new", category.NewCategory)
	route("/categories/edit/", category.UpdateCategory)
	route("/categories/delete/", category.DeleteCategory)

	route("/books/list", book.ListBooks)
	route("/books/get/", book.GetBook)
	route("/books/new", book.NewBook)
	route("/books/edit/", book.UpdateBook)
	route("/books/delete/", book.DeleteBook)

	log.Fatal(http.ListenAndServe(":2001", nil))
}

func route(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, middleware.PanicHandler(handler))
}
