package author

import (
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func NewAuthor(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	dto := getRequestBody(req)

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	authorEnt.CreateNewAuthor(ctx, dto)

	w.WriteHeader(204)
}

func getRequestBody(req *http.Request) *author.AuthorDto {
	var dto *author.AuthorDto

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dto)
	if err != nil {
		log.Fatalf("invalid request body :(")
	}

	return dto
}
