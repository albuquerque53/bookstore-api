package author

import (
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetAuthor(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	id := getAuthorId(req.URL.Path)

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	authors := authorEnt.GetAuthorById(ctx, id)

	resp, err := writer.ToJSON(authors)

	if err != nil {
		log.Fatalf("error during authors list: %s", err)
	}

	w.Write([]byte(resp))
}

func getAuthorId(urlPath string) int {
	strId := strings.TrimPrefix(urlPath, "/authors/get/")

	id, err := strconv.Atoi(strId)

	if err != nil {
		log.Fatalf("invalid author id :(")
	}

	return id
}
