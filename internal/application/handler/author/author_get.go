package author

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
)

func GetAuthor(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	id, err := helper.GetPathId(w, req.URL.Path)

	if err != nil {
		return
	}

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	author, err := authorEnt.GetAuthorById(ctx, id)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on author search", err)
		return
	}

	writer.SendResponse(w, 200, writer.JSONResponse{
		Message: "ok",
		Data:    author,
	})
}
