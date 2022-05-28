package author

import (
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
)

func ListAuthors(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	authors := authorEnt.GetAllAuthors(ctx)

	writer.SendResponse(w, 200, writer.JSONResponse{
		Message: "ok",
		Data:    authors,
	})
}
