package author

import (
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"log"
	"net/http"
)

func ListAuthors(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	authors := authorEnt.GetAllAuthors(ctx)

	resp, err := writer.ToJSON(authors)

	if err != nil {
		log.Fatalf("error during authors list: %s", err)
	}

	w.Write([]byte(resp))
}
