package author

import (
	"bookstoreapi/internal/application/helper"
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

	authors, err := authorEnt.GetAllAuthors(ctx)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on authors list", err)
		return
	}

	writer.SendResponse(w, 200, writer.JSONResponse{
		Message: "ok",
		Data:    authors,
	})
}
